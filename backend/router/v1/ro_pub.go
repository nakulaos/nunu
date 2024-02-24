/**
 ******************************************************************************
 * @file           : ro_pub.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/16
 ******************************************************************************
 */

package v1

import (
	"github.com/gin-gonic/gin"
	v1 "nunu/backend/api/v1"
)

type PubRouter struct {
}

func (p *PubRouter) InitPubRouter(Router *gin.RouterGroup, pubApi *v1.PubApi) {
	pubRouter := Router.Group("pub")
	{
		pubRouter.POST("/auth/register", pubApi.Register)
		pubRouter.GET("/captcha", pubApi.GetCaptcha)
		pubRouter.POST("/captcha/mail", pubApi.SendCaptchaWithMailReq)
		pubRouter.POST("/auth/login/username", pubApi.LoginWithUserName)
	}
}
