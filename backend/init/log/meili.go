/**
 ******************************************************************************
 * @file           : meili.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package log

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/sirupsen/logrus"
	"nunu/backend/init/conf"
	"sync"
)

var (
	_onceInitMeiliHook sync.Once
	_meiliLoggerHook   *meiliLogHook
)

type meiliLogData []map[string]any

type meiliLogHook struct {
	config    meilisearch.ClientConfig
	idxName   string
	addDocsCh chan *meiliLogData
}

func (h *meiliLogHook) Fire(entry *logrus.Entry) error {
	data := meiliLogData{{
		"id":      entry.Time.Unix(),
		"time":    entry.Time,
		"level":   entry.Level,
		"message": entry.Message,
		"data":    entry.Data,
	}}

	// 先尝试进log缓存，否则直接新开goroutine加文档
	select {
	case h.addDocsCh <- &data:
	default:
		go func(index *meilisearch.Index, item meiliLogData) {
			index.AddDocuments(item)
		}(h.index(), data)
	}

	return nil
}

func (h *meiliLogHook) handleAddDocs() {
	index := h.index()
	for item := range h.addDocsCh {
		index.AddDocuments(item)
	}
}

func (h *meiliLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *meiliLogHook) index() *meilisearch.Index {
	return meilisearch.NewClient(h.config).Index(h.idxName)
}

func initMeiliLogHook() {
	_onceInitMeiliHook.Do(func() {
		sysConfig := conf.GetConfig()
		_meiliLoggerHook = &meiliLogHook{
			config: meilisearch.ClientConfig{
				Host:   sysConfig.LoggerMeiliConf.Endpoint(),
				APIKey: sysConfig.LoggerMeiliConf.ApiKey,
			},
			idxName: sysConfig.LoggerMeiliConf.Index,
		}

		client := meilisearch.NewClient(_meiliLoggerHook.config)
		index := client.Index(_meiliLoggerHook.idxName)
		if _, err := index.FetchInfo(); err != nil {
			client.CreateIndex(&meilisearch.IndexConfig{
				Uid:        _meiliLoggerHook.idxName,
				PrimaryKey: "id",
			})
			sortableAttributes := []string{
				"time",
			}
			index.UpdateSortableAttributes(&sortableAttributes)
		}

		// 初始化addDocsCh
		_meiliLoggerHook.addDocsCh = make(chan *meiliLogData, sysConfig.LoggerMeiliConf.GetMaxLogBuffer())

		// 启动后台log工作者
		for minWork := sysConfig.LoggerMeiliConf.GetMinWork(); minWork > 0; minWork-- {
			go _meiliLoggerHook.handleAddDocs()
		}

	})

}

func GetMeiliLogHook() *meiliLogHook {
	initMeiliLogHook()
	return _meiliLoggerHook
}
