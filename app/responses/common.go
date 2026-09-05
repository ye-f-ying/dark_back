/*
 * @Date: 2026-05-31 18:00:59
 * @LastEditTime: 2026-05-31 18:01:00
 * @FilePath: /dark_back/app/responses/common.go
 * @Description:
 */
package responses

type MathCaptchaResponse struct {
	ImageBase64 string `json:"ImageBase64"`
	UUID        string `json:"uuid"`
}
