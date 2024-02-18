/**
 ******************************************************************************
 * @file           : err.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/14
 ******************************************************************************
 */

package constant

import "nunu/backend/pkg/xerror"

var (
	Success                     = xerror.NewError(0, "Success")
	ErrInternalServer           = xerror.NewError(10000, "ErrInternalServer")
	ErrBadRequest               = xerror.NewError(10001, "ErrBadRequest")
	ErrInvalidParams            = xerror.NewError(10002, "ErrInvalidParams")
	ErrNotFound                 = xerror.NewError(10003, "ErrNotFound")
	ErrUnauthorizedAuthNotExist = xerror.NewError(10004, "ErrUnauthorizedAuthNotExist")
	// ErrUnauthorizedAuthFailed 账号密码错误
	ErrUnauthorizedAuthFailed = xerror.NewError(10005, "ErrUnauthorizedAuthFailed")

	ErrUsernameHasExisted           = xerror.NewError(20001, "ErrUsernameHasExisted")
	ErrErrorCaptchaPassword         = xerror.NewError(20002, "ErrErrorCaptchaPassword")
	ErrTooManyPhoneCaptchaSend      = xerror.NewError(20003, "ErrTooManyPhoneCaptchaSend")
	ErrTooManyMailCaptchaSend       = xerror.NewError(20004, "ErrTooManyMailCaptchaSend")
	ErrExistedUserPhone             = xerror.NewError(20005, "ErrExistedUserPhone")
	ErrExistedUserMail              = xerror.NewError(20006, "ErrExistedUserMail")
	ErrErrorPhoneCaptcha            = xerror.NewError(20007, "ErrErrorPhoneCaptcha")
	ErrErrorMailCaptcha             = xerror.NewError(20008, "ErrErrorMailCaptcha")
	ErrEmailVerificationCodeExpired = xerror.NewError(20009, "ErrEmailVerificationCodeExpired")
	ErrMaxPhoneCaptchaUseTimes      = xerror.NewError(20010, "ErrMaxPhoneCaptchaUseTimes")
	ErrMaxMailCaptchaUseTimes       = xerror.NewError(20011, "ErrMaxMailCaptchaUseTimes")
	ErrUserRegisterFailed           = xerror.NewError(20012, "ErrUserRegisterFailed")
	ErrDisallowUserRegister         = xerror.NewError(20023, "ErrDisallowUserRegister")
)
