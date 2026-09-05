/*
 * @Date: 2026-05-31 14:43:51
 * @LastEditTime: 2026-05-31 17:40:29
 * @FilePath: /dark_back/app/handlers/common_handler.go
 * @Description:
 */
package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ye-f-ying/dark_back/app/services"
	"github.com/ye-f-ying/dark_back/internal/errors"
	"github.com/ye-f-ying/dark_back/internal/tools"
	"github.com/ye-f-ying/dark_pkg/pkg/db"
)

/**
 * @description: 图片验证码
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func MathCaptcha(ctx context.Context, c *app.RequestContext) {
	ip := tools.FormatIP(c.ClientIP())
	limitMax := 20

	allow, _, err := db.FrequencyCheckRedis(ctx, fmt.Sprintf("comm:math:cap:ip:%s", ip), limitMax, time.Second*60)
	if err != nil {
		adminErr := errors.NewAdminError(501, err, "")
		JSONAdminErrorResponse(c, ctx, adminErr)
		return
	}
	if !allow {
		JSONAdminErrorResponse(c, ctx, errors.NewAdminError(400, nil, "请求太过频繁！请等待%d分钟后再试！", 1))
		return
	}
	res, adminErr := services.GenerateMathCaptcha(ctx)
	if adminErr != nil {
		JSONAdminErrorResponse(c, ctx, adminErr)
		return
	}
	JSONData(c, res)
}
