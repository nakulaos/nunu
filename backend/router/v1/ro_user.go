/**
 ******************************************************************************
 * @file           : ro_user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/28
 ******************************************************************************
 */

package v1

import (
	"github.com/gin-gonic/gin"
	v1 "nunu/backend/api/v1"
	"nunu/backend/middleware"
)

type UserRouter struct {
}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup, userApi *v1.UserApi) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("/avatar", middleware.JWTAuth(), userApi.ChangeAvatar)
	}
}
