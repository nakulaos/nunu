/**
 ******************************************************************************
 * @file           : entry.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/27
 ******************************************************************************
 */

package middleware

import (
	"nunu/backend/init/gorm"
	"nunu/backend/persistence/jinzhu"
	"nunu/backend/repo"
)

var (
	_userManageRepo repo.IUserManageRepo
)

func init() {
	db := gorm.GetGormDB()
	_userManageRepo = jinzhu.NewUserManageRepo(db, jinzhu.NewUserMetricsRepo(db))
}
