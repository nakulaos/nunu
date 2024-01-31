/**
 ******************************************************************************
 * @file           : redis.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/30
 ******************************************************************************
 */

package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"nunu/backend/init/conf"
	"sync"
)

var (
	_rdb  *redis.Client
	_once sync.Once
)

func initRedisClient() {
	_once.Do(func() {

		redisConf := conf.GetConfig().RedisConf

		_rdb = redis.NewClient(&redis.Options{
			Addr:         redisConf.InitAddress[0],
			Password:     redisConf.Password,
			DB:           redisConf.SelectDB,
			WriteTimeout: redisConf.ConnWriteTimeout,
		})

		_, err := _rdb.Ping().Result()
		if err != nil {
			logrus.Errorf("redis连接失败:%s", redisConf.InitAddress[0])
		}

	})
}

func GetRedisClient() *redis.Client {
	initRedisClient()
	return _rdb
}
