/**
 ******************************************************************************
 * @file           : sqlx.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/29
 ******************************************************************************
 */

package sqlx

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"nunu/backend/init/conf"
	"sync"
)

var (
	_sqlxDB *sqlx.DB
	_once   sync.Once
)

func initSqlxDB() {
	sysConfig := conf.GetConfig()

	if sysConfig.MysqlConf.Enable {
		logrus.Debugln("use MySQL as relation database")
		_sqlxDB = sqlx.MustConnect("mysql", sysConfig.MysqlConf.Dsn())
	} else if sysConfig.PostgresConf.Enable {
		logrus.Debugln("use PostgresSQL as relation database")
		_sqlxDB = sqlx.MustConnect("postgres", sysConfig.PostgresConf.Dsn())
	} else if sysConfig.Sqlite3Conf.Enable {
		logrus.Debugln("use Sqlite3 as relation database")
		_sqlxDB = sqlx.MustConnect("sqlite3", sysConfig.Sqlite3Conf.Dsn("sqlite3"))
	} else {
		logrus.Debugln("use default MySQL as relation database")
		_sqlxDB = sqlx.MustConnect("MySQL", sysConfig.MysqlConf.Dsn())
	}

}

func GetSqlxDB() *sqlx.DB {
	_once.Do(func() {
		initSqlxDB()
	})
	return _sqlxDB
}
