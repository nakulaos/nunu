// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"nunu/backend/api/v1"
	"nunu/backend/init/gorm"
	"nunu/backend/init/mail"
	"nunu/backend/init/redis"
	"nunu/backend/persistence/cache"
	"nunu/backend/persistence/jinzhu"
	"nunu/backend/pkg/rand"
	"nunu/backend/service"
)

// Injectors from wire_injectors.go:

func CreateAdminApi() *v1.AdminApi {
	db := gorm.GetGormDB()
	iUserMetricsRepo := jinzhu.NewUserMetricsRepo(db)
	iUserManageRepo := jinzhu.NewUserManageRepo(db, iUserMetricsRepo)
	iAdminService := service.NewAdminService(iUserManageRepo)
	adminApi := v1.NewAdminApi(iAdminService)
	return adminApi
}

func CreatePubApi() *v1.PubApi {
	db := gorm.GetGormDB()
	iUserMetricsRepo := jinzhu.NewUserMetricsRepo(db)
	iUserManageRepo := jinzhu.NewUserManageRepo(db, iUserMetricsRepo)
	randRand := rand.NewRand()
	iMailVerifyServiceRepo := mail.GetMailVerify()
	iSecurityServiceRepo := jinzhu.NewSecurityServiceRepo(db, randRand, iMailVerifyServiceRepo)
	client := redis.GetRedisClient()
	iRedisCacheRepo := cache.NewRedisCacheRepo(client)
	iPubService := service.NewPubService(iUserManageRepo, iSecurityServiceRepo, iRedisCacheRepo)
	pubApi := v1.NewPubApi(iPubService)
	return pubApi
}

func CreateUserApi() *v1.UserApi {
	iUserService := service.NewUserService()
	userApi := v1.NewUserApi(iUserService)
	return userApi
}

// wire_injectors.go:

var allProviders = wire.NewSet(jinzhu.Set, v1.Set, service.Set, gorm.GetGormDB, mail.GetMailVerify, redis.GetRedisClient, rand.NewRand, cache.Set)
