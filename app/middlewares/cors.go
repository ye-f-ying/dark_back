/*
 * @Date: 2026-05-30 15:49:43
 * @LastEditTime: 2026-05-30 15:55:55
 * @FilePath: /dark_back/app/middlewares/cors.go
 * @Description:
 */
package middlewares

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"github.com/ye-f-ying/dark_back/internal/configs"
)

/**
 * @description: 跨域
 * @return {*}
 */
func CorsMiddleware() app.HandlerFunc {
	corsConfig := cors.Config{
		//AllowAllOrigins:  true, // 先允许所有
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	cfgs := configs.GetConfig()
	if len(cfgs.Hertz.AllowOrigins) <= 0 {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowAllOrigins = false
		corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, cfgs.Hertz.AllowOrigins...)
	}
	return cors.New(corsConfig)
}
