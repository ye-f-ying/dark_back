/*
 * @Date: 2026-05-31 17:56:30
 * @LastEditTime: 2026-05-31 18:25:51
 * @FilePath: /dark_back/app/requests/login.go
 * @Description:
 */
package requests

type LoginReq struct {
	Account string `json:"account" form:"account" validate:"required,min=2,max=20"`
	PWD     string `json:"pwd" form:"pwd" validate:"required,min=6,max=20"`
	Answer  int    `json:"answer" form:"answer" validate:"required"`
	UUID    string `json:"uuid" form:"uuid" validate:"required,min=1"`
}
