/*
 * @Date: 2026-05-31 14:57:21
 * @LastEditTime: 2026-05-31 15:04:03
 * @FilePath: /dark_back/app/constant/common.go
 * @Description:
 */
package constant

import (
	"fmt"
	"time"

	"github.com/ye-f-ying/dark_back/internal/tools"
)

const (
	COMMON_MATH_CAPTCHA_REDIS_KEY_PREFIX = "c:math:cap:"
	COMMON_MATH_CAPTCHA_REDIS_TIME       = time.Second * 60 * 5 // 5分钟有效
)

/**
 * @description:获取key
 * @param {string} uuid
 * @return {*}
 */
func GetCommonMathCaptchaRedisKey(uuid string) string {
	return fmt.Sprintf("%s%s", COMMON_MATH_CAPTCHA_REDIS_KEY_PREFIX, tools.ToLowerAndRemoveSpecialChars(uuid))
}
