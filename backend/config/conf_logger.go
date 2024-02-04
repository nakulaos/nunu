/**
 ******************************************************************************
 * @file           : conf_logger.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type LoggerConf struct {
	Level string
}

func (s *LoggerConf) GetLoggerLevel() logrus.Level {
	fmt.Println(s.Level)
	switch strings.ToLower(s.Level) {
	// 日志级别 panic|fatal|error|warn|info|debug|trace
	case "panic":
		return logrus.PanicLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "debug":
		return logrus.DebugLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "info":
		return logrus.InfoLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}
