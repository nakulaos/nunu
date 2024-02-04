/**
 ******************************************************************************
 * @file           : conf_web_server.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/4
 ******************************************************************************
 */

package config

import "time"

type WebServerConf struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
