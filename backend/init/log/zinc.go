/**
 ******************************************************************************
 * @file           : zinc.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package log

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
	"nunu/backend/init/conf"
	"sync"
	"time"
)

var (
	_zincLoggerHook   *zincLogHook
	_onceInitZincHook sync.Once
)

// 下面两个跟zinc有关
type zincLogData struct {
	Time    time.Time     `json:"time"`
	Level   logrus.Level  `json:"level"`
	Message string        `json:"message"`
	Data    logrus.Fields `json:"data"`
}

type zincLogIndex struct {
	Index map[string]string `json:"index"`
}

type zincLogHook struct {
	host     string // 上传的网址
	index    string // 上传的索引
	user     string // 上传的用户
	password string // 上传的密码
}

// Fire logrus的hook接口实现函数之一
func (h *zincLogHook) Fire(entry *logrus.Entry) error {
	// zinc 的api  上传数据要带索引
	// zinc 是使用 json 上传
	index := &zincLogIndex{Index: map[string]string{
		"_index": h.index,
	}}

	indexBytes, _ := json.Marshal(index)

	// 数据
	data := &zincLogData{
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
		Data:    entry.Data,
	}

	dataBytes, _ := json.Marshal(data)

	//
	logStr := string(indexBytes) + "\n" + string(dataBytes) + "\n"

	//resty是 Go 语言的一个 HTTP client 库。resty功能强大，特性丰富。
	//它支持几乎所有的 HTTP 方法（GET/POST/PUT/DELETE/OPTION/HEAD/PATCH等）
	//并提供了简单易用的 APIreturn nil
	client := resty.New()

	// 上传给zinc

	if _, err := client.SetDisableWarn(true).R().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(h.user, h.password).
		SetBody(logStr).Post(h.host); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

// Levels logrus的hook接口函数之一
func (h *zincLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func initZincHook() {
	_onceInitZincHook.Do(func() {
		sysConfig := conf.GetConfig()
		_zincLoggerHook = &zincLogHook{
			host: sysConfig.LoggerZincConf.EndPoint() + "/es/_bulk",
			// /es/_bulk是zinc 处理bulk 文件的一个接口
			index:    sysConfig.LoggerZincConf.Index,
			user:     sysConfig.LoggerZincConf.User,
			password: sysConfig.LoggerZincConf.Password,
		}
	})
}

func GetZincLogHook() *zincLogHook {
	initZincHook()
	return _zincLoggerHook
}
