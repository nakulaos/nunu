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
	"net/http"
	"nunu/backend/IService"
)

type AdminApi struct {
	Admin IService.IAdminService
}

func NewAdminApi(admin IService.IAdminService) *AdminApi {
	return &AdminApi{Admin: admin}
}

func (a *AdminApi) ChangeSiteInfo(c *gin.Context) {
	//req := &dto.ChangeUserStatusReq{}
	//err := c.ShouldBindJSON(req)
	//if err != nil {
	//
	//}
	//a.Admin.ChangeUserStatus(req)
	c.JSON(http.StatusOK, 0)
}
