/**
 ******************************************************************************
 * @file           : conf_redis.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/27
 ******************************************************************************
 */

package config

import "time"

type RedisConf struct {
	InitAddress      []string
	Username         string
	Password         string
	SelectDB         int
	ConnWriteTimeout time.Duration
}
