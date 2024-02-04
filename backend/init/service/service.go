/**
 ******************************************************************************
 * @file           : service.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package service

import (
	"nunu/backend/IService"
	"nunu/backend/init/repo"
	"nunu/backend/service"
	"sync"
)

var (
	_onceInitService sync.Once
	ImplAdminService IService.IAdminService
)

func InitService() {
	_onceInitService.Do(func() {
		ImplAdminService = service.NewAdminService(repo.ImplUserManageRepo)
	})
}
