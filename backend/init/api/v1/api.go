/**
 ******************************************************************************
 * @file           : api.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package v1

import (
	v1 "nunu/backend/api/v1"
	"nunu/backend/init/service"
	"sync"
)

var (
	ImplAdminApi *v1.AdminApi
	_onceInitApi sync.Once
)

func InitApi() {
	_onceInitApi.Do(func() {
		ImplAdminApi = v1.NewAdminApi(service.ImplAdminService)
	})
}
