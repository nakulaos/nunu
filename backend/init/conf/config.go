/**
 ******************************************************************************
 * @file           : config.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package conf

import (
	"fmt"
	viper2 "github.com/spf13/viper"
	"nunu/backend/config"
	"sync"
	"time"
)

var (
	_config *config.Config
	_once   sync.Once
	_viper  *viper2.Viper
)

func initConf() {
	_config = &config.Config{}

	_viper = viper2.New()
	_viper.SetConfigName("config")
	_viper.SetConfigType("yaml")
	_viper.AddConfigPath("../../etc")
	_viper.AddConfigPath("./backend/etc")

	err := _viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {              // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	objs := map[string]any{
		"MySQL":         &_config.MysqlConf,
		"Database":      &_config.DatabaseConf,
		"Postgres":      &_config.PostgresConf,
		"Sqlite3":       &_config.Sqlite3Conf,
		"Logger":        &_config.LoggerConf,
		"LoggerFile":    &_config.LoggerFileConf,
		"LoggerZinc":    &_config.LoggerZincConf,
		"LoggerMeili":   &_config.LoggerMeiliConf,
		"Sentry":        &_config.SentryConf,
		"Redis":         &_config.RedisConf,
		"Mail":          &_config.MailConf,
		"WebServer":     &_config.WebServerConf,
		"CustomProfile": &_config.CustomProfile,
	}

	for k, v := range objs {
		err = _viper.UnmarshalKey(k, v)
		if err != nil {
			panic(fmt.Errorf("Fatal error viper unmarshalkey:%s\n", err))
		}
	}

	_config.RedisConf.ConnWriteTimeout *= time.Second
}

func GetConfig() *config.Config {
	_once.Do(func() {
		initConf()
	})
	return _config
}

func GetViper() *viper2.Viper {
	_once.Do(func() {
		initConf()
	})
	return _viper
}
