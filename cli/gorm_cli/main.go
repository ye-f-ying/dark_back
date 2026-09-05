/*
 * @Date: 2026-06-15 18:06:51
 * @LastEditTime: 2026-06-16 15:14:20
 * @FilePath: /dark_back/cli/gorm_cli/main.go
 * @Description:
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ye-f-ying/dark_back/template/model_template"
	"github.com/ye-f-ying/dark_pkg/pkg/config"
	"github.com/ye-f-ying/dark_pkg/pkg/db"
	"github.com/ye-f-ying/dark_pkg/pkg/gorm_cli"
)

type Config struct {
	config.DefaultConfig `mapstructure:",squash"` // 需要扁平化
}

func main() {
	tableFlag := flag.String("table", "", "指定生成的表名，不指定则生成当前库所有表")
	outFlag := flag.String("out", "./models", "模型文件输出目录，默认 ./models")
	pkgName := flag.String("pkg", "models", "包名，默认 models")
	cfg, err := config.Init[*Config]()
	if err != nil {
		fmt.Println(err)
		return
	}
	conf, err := cfg.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)

	err = db.InitGormMYSQL()
	if err != nil {
		hlog.Errorf("init gorm db error:%v ", err)
		os.Exit(1)
		return
	}
	masterDB, err := db.GetGormDBMYSQL().DB()
	if err != nil {
		hlog.Errorf("init gorm db error:%v ", err)
		os.Exit(1)
		return
	}
	dbName := conf.GetMysql().DBName

	// 2. 创建输出目录（不存在则递归创建）
	outputDir := *outFlag
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ 创建输出目录失败: %v\n", err)
		return
	}

	var tables []string
	if *tableFlag != "" {
		tables = append(tables, *tableFlag)
		fmt.Printf("📌 指定生成单表: %s\n", *tableFlag)
	} else {
		tables, err = gorm_cli.GetAllTables(masterDB, dbName)
		if err != nil {
			fmt.Printf("❌ 获取库中所有表失败: %v\n", err)
			return
		}
		fmt.Printf("📌 检测到库中共有 %d 张表，开始批量生成\n", len(tables))
	}

	successCount := 0
	gorm_cli.GenerateBase(outputDir, *pkgName, model_template.ModelBaseTemplate)
	for _, tbl := range tables {
		meta, err := gorm_cli.GetTableMeta(masterDB, dbName, tbl)
		if err != nil {
			fmt.Printf("❌ 读取表 [%s] 结构失败: %v\n", tbl, err)
			continue
		}

		// 文件名 = 表名.go
		outPath := filepath.Join(outputDir, tbl+".go")
		err = gorm_cli.GenerateModel(meta, model_template.ModelTemplate, outPath, *pkgName)
		if err != nil {
			fmt.Printf("❌ 生成表 [%s] 模型失败: %v\n", tbl, err)
			continue
		}

		fmt.Printf("✅ 生成成功: %s\n", outPath)
		successCount++
	}

}
