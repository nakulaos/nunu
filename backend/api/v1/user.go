/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/28
 ******************************************************************************
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"nunu/backend/IService"
	"nunu/backend/api/v1/base"
	"nunu/backend/dto"
)

type UserApi struct {
	User IService.IUserService
}

func NewUserApi(user IService.IUserService) *UserApi {
	return &UserApi{User: user}
}

func (u *UserApi) ChangeAvatar(c *gin.Context) {
	req := new(dto.ChangeAvatarReq)
	if err := base.Bind(c, req); err != nil {
		base.Render(c, nil, err)
		return
	}
	err := u.User.ChangeAvatar(req)
	base.Render(c, nil, err)
}
