/**
 ******************************************************************************
 * @file           : conf_logger_zinc.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package config

type LoggerZincConf struct {
	Host     string
	Index    string
	User     string
	Password string
	Secure   bool // 采用http还是https
	Enable   bool
}

func endPoint(host string, secure bool) string {
	schema := "http"
	if secure {
		schema = "https"
	}
	return schema + "://" + host
}

func (s *LoggerZincConf) EndPoint() string {
	return endPoint(s.Host, s.Secure)
}
