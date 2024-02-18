/**
 ******************************************************************************
 * @file           : security.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/17
 ******************************************************************************
 */

package repo

import (
	"nunu/backend/model"
	"time"
)

// ISecurityServiceRepo  安全相关服务
type ISecurityServiceRepo interface {
	// GetLatestPhoneCaptcha 获取最新的手机验证码
	GetLatestPhoneCaptcha(phone string) (*model.Captcha, error)
	// UsePhoneCaptcha 更新短信验证码
	UsePhoneCaptcha(captcha *model.Captcha) error
	// SendPhoneCaptcha 发送短信验证码
	SendPhoneCaptcha(phone string) error

	// GetLatestMailCaptcha 获取最新的手机验证码
	GetLatestMailCaptcha(mail string) (*model.Captcha, error)
	// UseMailCaptcha 更新短信验证码
	UseMailCaptcha(captcha *model.Captcha) error
	// SendMailCaptcha 发送短信验证码
	SendMailCaptcha(mail string) error
}

// IMailVerifyServiceRepo MailVerifyService 邮件验证服务
type IMailVerifyServiceRepo interface {
	// SendMailCaptcha  发送邮件验证码
	SendMailCaptcha(mail string, captcha string, expire time.Duration) error
}
