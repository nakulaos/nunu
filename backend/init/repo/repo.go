/**
 ******************************************************************************
 * @file           : repo.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"nunu/backend/init/conf"
	gorm2 "nunu/backend/init/gorm"
	"nunu/backend/init/log"
	"nunu/backend/persistence/jinzhu"
	"nunu/backend/repo"
	"sync"
)

var (
	ImplUserManageRepo repo.IUserManageRepo
	_onceInitial       sync.Once
)

func InitRepo() {
	_onceInitial.Do(func() {
		databaseConf := conf.GetConfig().DatabaseConf
		log.InitLogger()
		if databaseConf.Orm == "Gorm" {
			gormDB := gorm2.GetGormDB()
			initRepoWithGormDB(gormDB)
		} else {
			logrus.Warnln("This feature is not supported in the current version!")
		}
	})
}

func initRepoWithGormDB(db *gorm.DB) {
	// TODO: 等待补充
	ImplUserManageRepo = jinzhu.NewUserManageRepo(db)
}
