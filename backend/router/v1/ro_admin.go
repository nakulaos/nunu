/**
 ******************************************************************************
 * @file           : ro_admin.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package v1

import (
	"github.com/gin-gonic/gin"
	v1 "nunu/backend/api/v1"
)

type AdminRouter struct {
}

func (a *AdminRouter) InitAdminRouter(Router *gin.RouterGroup, adminApi *v1.AdminApi) {
	adminRouter := Router.Group("admin")
	{
		adminRouter.GET("/site/status", adminApi.ChangeSiteInfo)
	}
}
