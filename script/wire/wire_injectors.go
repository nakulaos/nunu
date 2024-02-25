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
	"nunu/backend/init/mail"
	"nunu/backend/init/redis"
	"nunu/backend/persistence/cache"
	"nunu/backend/persistence/jinzhu"
	"nunu/backend/pkg/rand"
	"nunu/backend/service"
)

var allProviders = wire.NewSet(
	jinzhu.Set,
	v1.Set,
	service.Set,
	gorm.GetGormDB,
	mail.GetMailVerify,
	redis.GetRedisClient,
	rand.NewRand,
	cache.Set,
)

func CreateAdminApi() *v1.AdminApi {
	wire.Build(
		allProviders)
	return new(v1.AdminApi)
}

func CreatePubApi() *v1.PubApi {
	wire.Build(
		allProviders)
	return new(v1.PubApi)
}
