/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/28
 ******************************************************************************
 */

package service

import (
	"nunu/backend/IService"
	"nunu/backend/dto"
	"nunu/backend/pkg/xerror"
)

var _ (IService.IUserService) = (*UserService)(nil)

type UserService struct {
}

func NewUserService() IService.IUserService {
	return &UserService{}
}

func (u *UserService) ChangeAvatar(req *dto.ChangeAvatarReq) xerror.Error {
	return nil
}
