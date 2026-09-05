/*
 * @Date: 2026-05-31 17:48:29
 * @LastEditTime: 2026-05-31 18:05:00
 * @FilePath: /dark_back/internal/tools/hertz.go
 * @Description:
 */
package tools

import (
	"reflect"
	"strings"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-playground/validator/v10"
)

/**
 * @description: 获取前端所有请求数据
 * @param {*app.RequestContext} c
 * @param {any} obj
 * @return {*}
 */
func RequestAll(c *app.RequestContext, obj any) error {
	contentType := strings.ToLower(string(c.ContentType()))
	// 判断是否为 JSON（兼容带 charset 的情况）
	if strings.HasPrefix(contentType, consts.MIMEApplicationJSON) {
		return c.BindJSON(obj)
	}
	return c.Bind(obj)
}

var validate *validator.Validate
var once sync.Once

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("json")
		})
	})
	return validate
}

// Validate 验证结构体
func Validate(obj any) error {
	return GetValidator().Struct(obj)
}
