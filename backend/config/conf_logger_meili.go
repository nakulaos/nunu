/**
 ******************************************************************************
 * @file           : conf_logger_meili.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package config

type LoggerMeiliConf struct {
	Host         string
	Index        string
	ApiKey       string
	Secure       bool
	MaxLogBuffer int
	MinWorker    int
	Enable       bool
}

func (s *LoggerMeiliConf) Endpoint() string {
	return endPoint(s.Host, s.Secure)
}

func (s *LoggerMeiliConf) GetMinWork() int {
	if s.MinWorker < 5 {
		return 5
	} else if s.MinWorker > 100 {
		return 100
	}
	return s.MinWorker
}

func (s *LoggerMeiliConf) GetMaxLogBuffer() int {
	if s.MaxLogBuffer < 10 {
		return 10
	} else if s.MaxLogBuffer > 1000 {
		return 1000
	}
	return s.MaxLogBuffer
}
