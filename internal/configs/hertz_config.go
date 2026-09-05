/*
 * @Date: 2026-05-30 15:29:32
 * @LastEditTime: 2026-05-30 18:14:24
 * @FilePath: /dark_back/internal/configs/hertz_config.go
 * @Description:
 */
package configs

type HertzConfig struct {
	Port         int      `mapstructure:"port"  default:"12000"` // Hertz监听的端口
	AllowOrigins []string `mapstructure:"allow_origins"`
}
