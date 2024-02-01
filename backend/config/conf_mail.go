/**
 ******************************************************************************
 * @file           : conf_mail.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/27
 ******************************************************************************
 */

package config

type MailConf struct {
	HttpHost         string
	HttpPort         int
	User             string
	Password         string
	DefaultFromEmail string
	UseSSL           bool
	UserTLS          bool
}
