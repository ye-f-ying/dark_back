/*
 * @Date: 2026-05-31 14:44:55
 * @LastEditTime: 2026-05-31 18:01:15
 * @FilePath: /dark_back/app/services/common_service.go
 * @Description:
 */
package services

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/ye-f-ying/dark_back/app/constant"
	"github.com/ye-f-ying/dark_back/app/responses"
	"github.com/ye-f-ying/dark_back/internal/errors"
	"github.com/ye-f-ying/dark_pkg/pkg/db"
	"github.com/ye-f-ying/dark_pkg/pkg/utils"
)

/**
 * @description:生成数学验证码
 * @param {context.Context} ctx
 * @return {*}
 */
func GenerateMathCaptcha(ctx context.Context) (*responses.MathCaptchaResponse, *errors.AdminError) {
	mc, err := utils.GenerateMathCaptcha()
	if err != nil {
		return nil, errors.NewAdminError(1000, err, "验证码生成异常")
	}

	if mc == nil || mc.ImageBase64 == "" || mc.Answer == "" {
		return nil, errors.NewAdminError(1000, nil, "验证码生成工具异常！")
	}

	uid := uuid.NewString()

	err = db.GetRedisClient().
		Set(ctx, constant.GetCommonMathCaptchaRedisKey(uid), mc.Answer, constant.COMMON_MATH_CAPTCHA_REDIS_TIME).
		Err()
	if err != nil {
		hlog.Errorf("[/dark_back/app/services/common_service.go->GenerateMathCaptcha]出现redis set操作异常！信息：%v", err)
		return nil, errors.NewAdminError(501, err, "")
	}

	return &responses.MathCaptchaResponse{
		ImageBase64: mc.ImageBase64,
		UUID:        uid,
	}, nil
}

/**
 * @description: 验证验证码
 * @param {context.Context} ctx
 * @param {*} UUID
 * @param {string} answer
 * @return {*}
 */
func VerifyCaptcha(ctx context.Context, UUID, answer string) (bool, *errors.AdminError) {
	if answer == "" {
		return false, nil
	}
	key := constant.GetCommonMathCaptchaRedisKey(UUID)
	result, err := db.GetRedisClient().Get(ctx, key).
		Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		hlog.Errorf("[/dark_back/app/services/common_service.go->VerifyCaptcha]出现redis get操作异常！信息：%v", err)
		return false, errors.NewAdminError(501, err, "")
	}

	if result != answer {
		return false, nil
	}
	_ = db.GetRedisClient().Del(ctx, key).Err()
	return true, nil
}
