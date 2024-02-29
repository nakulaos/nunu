/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/26
 ******************************************************************************
 */

package IService

import (
	"nunu/backend/dto"
	"nunu/backend/pkg/xerror"
)

type IUserService interface {
	ChangeAvatar(req *dto.ChangeAvatarReq) xerror.Error
}
