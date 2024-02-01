/**
 ******************************************************************************
 * @file           : migration.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/1
 ******************************************************************************
 */

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"nunu/backend/init/conf"
	"nunu/backend/init/gorm"
	"nunu/backend/init/log"
	"nunu/backend/repo/model"
)

func initMigrationCommand() {
	_commandMigration := &cobra.Command{
		Use:   "migration",
		Short: "the command is used for migration database source",
		Long:  "the command is used for migration database source",
		Run:   migrationCommandHandle,
	}

	rootCmd.AddCommand(_commandMigration)
}

func migrationCommandHandle(_cmd *cobra.Command, _args []string) {
	dataBaseConf := conf.GetConfig().DatabaseConf

	log.InitLogger()

	if dataBaseConf.Orm == "Gorm" {

		db := gorm.GetGormDB()

		logrus.Debugln("Database is being migrated...")
		err := db.AutoMigrate(
			//TODO: 等待补充...
			&model.User{},
		)
		logrus.Errorf("An error occurred while migrating the database:%s", err)

	} else {
		logrus.Warnln("This feature is not supported at the moment!")
	}
}
