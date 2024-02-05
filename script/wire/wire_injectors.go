//go:build wireinject
// +build wireinject

/**
 ******************************************************************************
 * @file           : wire_injectors.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/5
 ******************************************************************************
 */

package wire

import (
	"github.com/google/wire"
	v1 "nunu/backend/api/v1"
	"nunu/backend/init/gorm"
	"nunu/backend/persistence/jinzhu"
	"nunu/backend/service"
)

var allProviders = wire.NewSet(
	jinzhu.Set,
	v1.Set,
	service.Set,
	gorm.GetGormDB,
)

func CreateAdminApi() *v1.AdminApi {
	wire.Build(
		allProviders)
	return new(v1.AdminApi)
}
