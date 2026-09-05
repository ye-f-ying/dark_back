/*
 * @Date: 2026-05-31 14:26:06
 * @LastEditTime: 2026-05-31 14:26:08
 * @FilePath: /dark_back/internal/locales/locales.go
 * @Description:
 */
package locales

import "context"

type langContextKey string

const (
	contextKeyLang langContextKey = "lang" // 语言
)

/**
 * @description:获取语言
 * @param {context.Context} ctx
 * @return {*}
 */
func GetLangFromCtx(ctx context.Context) string {
	lang, ok := ctx.Value(contextKeyLang).(string)
	if !ok {
		return "zh-CN" // 默认中文
	}
	return lang
}
