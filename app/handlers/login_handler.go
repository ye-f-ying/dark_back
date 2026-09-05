/*
 * @Date: 2026-05-31 17:40:56
 * @LastEditTime: 2026-05-31 18:17:39
 * @FilePath: /dark_back/app/handlers/login_handler.go
 * @Description:
 */
package handlers

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ye-f-ying/dark_back/app/requests"
	"github.com/ye-f-ying/dark_back/app/services"
	"github.com/ye-f-ying/dark_back/internal/errors"
)

/**
 * @description: admin 账号登录
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func AdminLogin(ctx context.Context, c *app.RequestContext) {
	var login requests.LoginReq
	adminErr := RequestAll(c, &login)
	if adminErr != nil {
		JSONAdminErrorResponse(c, ctx, adminErr)
		return
	}

	// 验证图片验证码
	result, adminErr := services.VerifyCaptcha(ctx, login.UUID, strconv.Itoa(login.Answer))
	if adminErr != nil {
		JSONAdminErrorResponse(c, ctx, adminErr)
		return
	}

	if !result {
		JSONAdminErrorResponse(c, ctx, errors.NewAdminError(1001, nil, "验证码错误！"))
		return
	}

	//

}
