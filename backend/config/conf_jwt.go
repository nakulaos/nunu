/**
 ******************************************************************************
 * @file           : conf_jwt.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/24
 ******************************************************************************
 */

package config

import "time"

type JwtConf struct {
	Secret string
	Issuer string
	Expire time.Duration
}
