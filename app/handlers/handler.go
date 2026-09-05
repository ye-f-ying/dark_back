/*
 * @Date: 2026-05-31 14:00:56
 * @LastEditTime: 2026-05-31 18:28:25
 * @FilePath: /dark_back/app/handlers/handler.go
 * @Description:
 */
package handlers

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ye-f-ying/dark_back/internal/errors"
	"github.com/ye-f-ying/dark_back/internal/tools"
)

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据
	Error   string      `json:"error"`   // 错误信息
}

/**
 * @description: admin 错误转换为返回前端数据
 * @param {context.Context} ctx
 * @param {errors.AdminError} err
 * @param {interface{}} data
 * @return {*}
 */
func AdminErrorToResponse(ctx context.Context, err errors.AdminError, data interface{}) Response {
	if err.Code == errors.NORMAL_ERROR_CODE {
		return Response{
			Code:    errors.NORMAL_ERROR_CODE,
			Message: "success!", // 服务器错误翻译为前端提示错误
			Error:   "",         //原始错误
			Data:    data,
		}
	}
	return Response{
		Code:    err.Code,
		Message: errors.GetErrorMsg(ctx, err.Code, err.Args...), // 服务器错误翻译为前端提示错误
		Error:   err.Error(),                                    //原始错误
	}
}

/**
 * @description: json 返回格式统一
 * @param {*app.RequestContext} c
 * @param {Response} res
 * @return {*}
 */
func JSONResponse(c *app.RequestContext, res Response) {
	c.JSON(http.StatusOK, res)
}

/**
 * @description:json 数据返回
 * @param {*app.RequestContext} c
 * @param {interface{}} data
 * @return {*}
 */
func JSONData(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    errors.NORMAL_ERROR_CODE,
		Message: "success!", // 服务器错误翻译为前端提示错误
		Error:   "",         //原始错误
		Data:    data,
	})
}

/**
 * @description: json 返回错误格式统一
 * @param {*app.RequestContext} c
 * @param {context.Context} ctx
 * @param {errors.AdminError} err
 * @return {*}
 */
func JSONAdminErrorResponse(c *app.RequestContext, ctx context.Context, err *errors.AdminError) {
	if err == nil {
		c.JSON(http.StatusOK, AdminErrorToResponse(ctx, *errors.NewAdminError(500, nil, ""), nil))
		return
	}
	c.JSON(http.StatusOK, AdminErrorToResponse(ctx, *err, nil))
}

/**
 * @description: 获取所有参数并自动验证参数
 * @param {*app.RequestContext} c
 * @param {any} obj
 * @return {*}
 */
func RequestAll(c *app.RequestContext, obj any) *errors.AdminError {
	if obj == nil {
		return errors.NewAdminError(503, nil, "传入的参数不能是nil！")
	}
	err := tools.RequestAll(c, obj)
	if err != nil {
		return errors.NewAdminError(504, err, "解析/读取参数失败！")
	}
	err = tools.Validate(obj)
	if err != nil {
		return errors.NewAdminError(201, err, "参数验证失败！")
	}
	return nil
}
