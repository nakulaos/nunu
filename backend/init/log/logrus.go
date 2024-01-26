/**
 ******************************************************************************
 * @file           : logrus.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package log

import (
	"github.com/sirupsen/logrus"
	"nunu/backend/init/conf"
)

func InitLogger() {
	sysConfig := conf.GetConfig()
	// 设置格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志等级
	logrus.SetLevel(sysConfig.LoggerConf.GetLoggerLevel())

	// 配置日志源
	if sysConfig.LoggerFileConf.Enable {
		// 配置是同LoggerFile
		out := GetFileLogger()
		logrus.SetOutput(out)
	}

	if sysConfig.LoggerZincConf.Enable {
		// 配置zinc
		hook := GetZincLogHook()
		logrus.AddHook(hook)
	}

	if sysConfig.LoggerMeiliConf.Enable {
		// 配置meili
		hook := GetMeiliLogHook()
		logrus.AddHook(hook)
	}

}
