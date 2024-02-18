/**
 ******************************************************************************
 * @file           : pub.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package IService

import (
	"nunu/backend/dto"
	"nunu/backend/pkg/xerror"
)

type IPubService interface {
	SendCaptchaWithMailReq(req *dto.SendCaptchaWithMailReq) xerror.Error
	SendCaptchaWithPhoneReq(req *dto.SendCaptchaWithPhoneReq) xerror.Error
	GetCaptcha() (*dto.GetCaptchaResp, xerror.Error)
	Register(req *dto.RegisterReq) (*dto.RegisterResp, xerror.Error)
	Login(req *dto.LoginReq) (*dto.LoginResp, xerror.Error)
}
