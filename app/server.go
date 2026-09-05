/*
 * @Date: 2026-05-22 15:41:14
 * @LastEditTime: 2026-06-15 18:17:04
 * @FilePath: /dark_back/app/server.go
 * @Description:
 */
package app

import (
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ye-f-ying/dark_back/app/routers"
	"github.com/ye-f-ying/dark_back/internal/configs"
	"github.com/ye-f-ying/dark_pkg/pkg/db"
	"github.com/ye-f-ying/dark_pkg/pkg/zap"
)

func Run() {
	// 读取配置文件
	cfg := configs.GetConfig()
	// 初始化zap 日志
	zap.InitZap(cfg)
	// 初始化redis
	err := db.InitRedis()
	if err != nil {
		hlog.Errorf("init redis error:%v ", err)
		os.Exit(1)
		return
	}
	// 初始化mysql
	err = db.InitGormMYSQL()
	if err != nil {
		hlog.Errorf("init grom db error:%v ", err)
		os.Exit(1)
		return
	}

	// hertz 框架
	h := server.Default(server.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", cfg.Hertz.Port)))
	// 加载路由
	routers.Router(h)
	h.Spin()
}
