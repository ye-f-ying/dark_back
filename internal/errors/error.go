/*
 * @Date: 2026-05-31 14:01:48
 * @LastEditTime: 2026-05-31 16:11:32
 * @FilePath: /dark_back/internal/errors/error.go
 * @Description:
 */
package errors

import (
	"fmt"
)

const NORMAL_ERROR_CODE = 200

type AdminError struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Err     error  `json:"error"`   // 错误信息
	Args    []any
}

func (m *AdminError) Error() string {
	if m.Err == nil {
		return m.Message
	}
	return m.Err.Error()
}

/**
 * @description: ADMIN 错误
 * @param {context.Context} ctx
 * @param {int} code
 * @param {error} err
 * @param {string} msg
 * @param {...interface{}} errs
 * @return {*}
 */
func NewAdminError(code int, err error, msg string, errs ...interface{}) *AdminError {
	return &AdminError{
		Code:    code,
		Err:     err,
		Message: fmt.Sprintf(msg, errs...),
		Args:    errs,
	}
}
