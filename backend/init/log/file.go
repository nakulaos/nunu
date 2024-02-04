/**
 ******************************************************************************
 * @file           : file.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"nunu/backend/init/conf"
	"sync"
)

var (
	_onceInitFile sync.Once
	loggerFile    io.Writer
)

func initFileLogger() {
	_onceInitFile.Do(func() {
		sysConfig := conf.GetConfig()
		loggerFile = &lumberjack.Logger{
			Filename:  sysConfig.LoggerFileConf.SavePath + "/" + sysConfig.LoggerFileConf.FileName + sysConfig.LoggerFileConf.FileExt,
			MaxSize:   500,
			MaxAge:    28,
			LocalTime: true,
		}
	})
}

func GetFileLogger() io.Writer {
	initFileLogger()
	return loggerFile
}
