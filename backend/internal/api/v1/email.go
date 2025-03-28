package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/cache"
	"github.com/joey17520/ailiaili/internal/domain/dto"
	"github.com/joey17520/ailiaili/internal/global"
	"github.com/joey17520/ailiaili/internal/resp"
	"github.com/joey17520/ailiaili/pkg/mail"
	"github.com/joey17520/ailiaili/utils"
	"go.uber.org/zap"
)

func SendRegisterEmailCode(ctx *gin.Context) {
	// 获取参数
	var sendEmailReq dto.SendEmailReq
	if err := ctx.Bind(&sendEmailReq); err != nil {
		resp.FailWithMessage(ctx, "请求参数有误")
		return
	}

	// 参数校验
	if !utils.VerifyEmail(sendEmailReq.Email) {
		resp.FailWithMessage(ctx, "邮箱格式错误")
		return
	}

	if utils.VerifyStringLength(sendEmailReq.CaptchaId, "=", 0) {
		captchaId := cache.CreateCaptchaStatus()
		resp.Result(ctx, -1, gin.H{"captchaId": captchaId}, "需要人机验证")
		return
	}

	// 如果未进行人机验证
	if cache.GetCaptchaStatus(sendEmailReq.CaptchaId) == global.CAPTCHA_STATUS_ABSENT {
		captchaId := cache.CreateCaptchaStatus()
		resp.Result(ctx, -1, gin.H{"captchaId": captchaId}, "需要人机验证")
		return
	}

	// 生成code
	code := utils.GenerateNumberCode(6)

	// 发送code
	if global.Config.Mail.Debug {
		// 验证码debug模式不发送邮件
		zap.L().Debug("邮箱:" + sendEmailReq.Email + ",验证码:" + code)
	} else {
		if err := mail.SendCaptcha(sendEmailReq.Email, code); err != nil {
			utils.ErrorLog("邮箱验证码发送失败", "email", err.Error())
			resp.FailWithMessage(ctx, "邮箱验证码发送失败")
			return
		}
	}

	// code放入缓存
	cache.SetEmailCode(sendEmailReq.Email, code)
	// 删除人机验证状态
	cache.DelCaptchaStatus(sendEmailReq.CaptchaId)

	resp.Ok(ctx)
}
