/**
 ******************************************************************************
 * @file           : ro_admin.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package service

import (
	"nunu/backend/IService"
	"nunu/backend/dto"
	"nunu/backend/pkg/xerror"
	"nunu/backend/repo"
)

var _ IService.IAdminService = (*AdminService)(nil)

type AdminService struct {
	UserManageRepo repo.IUserManageRepo
}

func NewAdminService(userManageRepo repo.IUserManageRepo) IService.IAdminService {
	return &AdminService{UserManageRepo: userManageRepo}
}

func (a *AdminService) ChangeUserStatus(req *dto.ChangeUserStatusReq) {
	// TODO implement me
	panic("implement me")
}

func (a *AdminService) SiteInfo(req *dto.SiteInfoReq) (resp *dto.SiteInfoResp, err xerror.Error) {
	// TODO implement me
	panic("implement me")
}
