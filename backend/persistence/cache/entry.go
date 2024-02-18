/**
 ******************************************************************************
 * @file           : entry.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/18
 ******************************************************************************
 */

package cache

import "nunu/backend/constant"

var (
	_imgCaptchaKey   string
	_mailCaptchaKey  string
	_phoneCaptchaKey string
)

func init() {
	_imgCaptchaKey = constant.ImgCaptchaKey
	_mailCaptchaKey = constant.MailCaptchaKey
	_phoneCaptchaKey = constant.PhoneCaptchaKey
}
