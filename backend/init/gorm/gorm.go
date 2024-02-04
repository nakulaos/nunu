/**
 ******************************************************************************
 * @file           : jinzhu.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package gorm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"nunu/backend/init/conf"
	"sync"
	"time"
)

var (
	_once   sync.Once
	_gormdb *gorm.DB
)

func initGormDB() {
	sysConfig := conf.GetConfig()
	newLogger := logger.New(
		logrus.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  sysConfig.DatabaseConf.GetLogLevel(),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		})

	config := &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   sysConfig.DatabaseConf.TablePrefix,
			SingularTable: true,
		},
	}

	plugin := dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(sysConfig.MysqlConf.MaxIdleConns).
		SetMaxOpenConns(sysConfig.MysqlConf.MaxOpenConns)

	if sysConfig.MysqlConf.Enable {
		logrus.Debugln("use mysql as db")
		var err error
		_gormdb, err = gorm.Open(mysql.Open(sysConfig.MysqlConf.Dsn()), config)
		if err == nil {
			_gormdb.Use(plugin)
		} else {
			panic(fmt.Sprintf("failure to using mysql :%s", err))
		}
	} else if sysConfig.PostgresConf.Enable {
		logrus.Debugln("use postgres mysql as db")
		var err error
		_gormdb, err = gorm.Open(postgres.Open(sysConfig.PostgresConf.Dsn()), config)
		if err != nil {
			panic(fmt.Sprintf("failure to use postgres : %s", err))
		}
	} else if sysConfig.Sqlite3Conf.Enable {
		logrus.Debugln("use sqlite3 as db")
		var err error
		_gormdb, err = gorm.Open(sqlite.Open(sysConfig.Sqlite3Conf.Dsn("sqlite3")), config)
		if err != nil {
			panic(fmt.Sprintf("failure to use sqlite3 : %s ", err))
		}
	} else {
		logrus.Debugln("use default of mysql as db")
		var err error
		_gormdb, err = gorm.Open(mysql.Open(sysConfig.MysqlConf.Dsn()), config)
		if err == nil {
			_gormdb.Use(plugin)
		} else {
			panic(fmt.Sprintf("failure to using mysql :%s", err))
		}
	}
}

func GetGormDB() *gorm.DB {
	_once.Do(func() {
		initGormDB()
	})
	return _gormdb
}
