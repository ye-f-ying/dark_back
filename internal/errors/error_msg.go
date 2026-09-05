/*
 * @Date: 2026-05-31 14:14:43
 * @LastEditTime: 2026-05-31 18:29:01
 * @FilePath: /dark_back/internal/errors/error_msg.go
 * @Description:
 */
package errors

import (
	"context"
	"fmt"

	"github.com/ye-f-ying/dark_back/internal/locales"
)

var errorMsg = map[string]map[int]string{
	// 中文
	"zh-CN": {
		200: "成功！",
		201: "参数验证失败！",

		400: "请求太过频繁！请等待%d分钟后再试！",

		501: "内部服务器异常！", //redis 相关
		502: "内部服务器异常！", //mysql 相关
		503: "传入的参数不能是nil！",
		504: "解析/读取参数失败！",

		1000: "验证码生成失败！",
		1001: "验证码错误！",
		1003: "用户不存在",
	},
}

/**
 * @description: 获取错误信息
 * @param {context.Context} ctx
 * @param {int} code
 * @param {...interface{}} msg
 * @return {*}
 */
func GetErrorMsg(ctx context.Context, code int, msg ...interface{}) string {
	lang := locales.GetLangFromCtx(ctx)
	if lang == "" {
		lang = "zh-CN"
	}
	if errorMsg[lang] == nil || errorMsg[lang][code] == "" {
		if len(msg) <= 0 {
			return "unknown error"
		}
		return fmt.Sprint(msg...)
	}
	return fmt.Sprintf(errorMsg[lang][code], msg...)
}
