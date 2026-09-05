/*
 * @Date: 2026-05-30 15:24:45
 * @LastEditTime: 2026-05-30 18:14:34
 * @FilePath: /dark_back/internal/configs/config.go
 * @Description:
 */
package configs

import (
	"os"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ye-f-ying/dark_pkg/pkg/config"
)

type Config struct {
	config.DefaultConfig `mapstructure:",squash"` // 需要扁平化
	Hertz                HertzConfig              `mapstructure:"hertz"` // 分布式ID
}

var cfg *Config
var cfgOnce sync.Once

/**
 * @description: 获取配置文件
 * @return {*}
 */
func GetConfig() *Config {
	cfgOnce.Do(func() {
		cfgTemp, err := config.Init[*Config]()
		if err != nil {
			hlog.Errorf("config init error:%v", err)
			os.Exit(1)
			return
		}
		conf, err := cfgTemp.GetConfig()
		if err != nil || conf == nil {
			if err != nil {
				hlog.Errorf("config error:%v", err)
			} else {
				hlog.Errorf("config is nil")
			}
			os.Exit(1)
		}
		cfg = conf
	})
	return cfg
}
