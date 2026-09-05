/*
 * @Date: 2026-05-30 15:59:48
 * @LastEditTime: 2026-05-31 18:18:41
 * @FilePath: /dark_back/app/routers/admin.go
 * @Description:
 */
package routers

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/ye-f-ying/dark_back/app/handlers"
)

func adminRouters(r *route.RouterGroup) {
	r.GET("/common/math_captcha", handlers.MathCaptcha)

	r.POST("/account/login", handlers.AdminLogin)
}
