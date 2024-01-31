/**
 ******************************************************************************
 * @file           : config.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/23
 ******************************************************************************
 */

package config

type Config struct {
	MysqlConf       MysqlConf
	Database        DatabaseConf
	PostgresConf    PostgresConf
	Sqlite3Conf     Sqlite3Conf
	LoggerConf      LoggerConf
	LoggerFileConf  LoggerFileConf
	LoggerZincConf  LoggerZincConf
	LoggerMeiliConf LoggerMeiliConf
	SentryConf      SentryConf
	RedisConf       RedisConf
}
