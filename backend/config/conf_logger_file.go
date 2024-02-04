/**
 ******************************************************************************
 * @file           : conf_logger_file.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package config

type LoggerFileConf struct {
	SavePath string
	FileName string
	FileExt  string
	Enable   bool
}
