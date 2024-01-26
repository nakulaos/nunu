/**
 ******************************************************************************
 * @file           : conf_database.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/25
 ******************************************************************************
 */

package config

import (
	"gorm.io/gorm/logger"
	"strings"
)

type DatabaseConf struct {
	TablePrefix string
	LogLevel    string
}

func (s *DatabaseConf) GetLogLevel() logger.LogLevel {
	switch strings.ToLower(s.LogLevel) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Error
	}
}
