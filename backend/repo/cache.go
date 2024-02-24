/**
 ******************************************************************************
 * @file           : cache.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/23
 ******************************************************************************
 */

package repo

// RedisCache memory cache by Redis
type IRedisCacheRepo interface {
	// SetPushToSearchJob 设置推送搜索作业
	SetPushToSearchJob() error
	// DelPushToSearchJob 删除推送搜索作业
	DelPushToSearchJob() error
	// SetImgCaptcha 设置图片验证码
	SetImgCaptcha(id string, value string) error
	// GetImgCaptcha 获取图片验证码
	GetImgCaptcha(id string) (string, error)
	// DelImgCaptcha 删除图片验证码
	DelImgCaptcha(id string) error
	// GetCountPhoneCaptcha  获取验证码频次限制
	GetCountPhoneCaptcha(phone string) (int64, error)
	// IncrCountPhoneCaptcha  增加验证码频次
	IncrCountPhoneCaptcha(phone string) error

	// GetCountMailCaptcha  获取验证码频次限制
	GetCountMailCaptcha(mail string) (int64, error)
	// IncrCountMailCaptcha 增加验证码频次
	IncrCountMailCaptcha(mail string) error

	// GetCountLoginErr 获取登录错误次数
	GetCountLoginErr(id uint64) (int64, error)
	// DelCountLoginErr 清空登录错误次数
	DelCountLoginErr(id uint64) error
	// IncrCountLoginErr 增加登录错误次数
	IncrCountLoginErr(id uint64) error
	// SetRechargeStatus 设置付款状态
	SetRechargeStatus(tradeNo string) error
	// DelRechargeStatus 删除付款状态
	DelRechargeStatus(tradeNo string) error
}
