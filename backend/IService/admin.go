/**
 ******************************************************************************
 * @file           : ro_admin.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package IService

import (
	"nunu/backend/dto"
	"nunu/backend/pkg/xerror"
)

type IAdminService interface {
	ChangeUserStatus(req *dto.ChangeUserStatusReq)
	SiteInfo(req *dto.SiteInfoReq) (resp *dto.SiteInfoResp, err xerror.Error)
}
