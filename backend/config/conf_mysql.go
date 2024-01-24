/**
 ******************************************************************************
 * @file           : conf_mysql.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/24
 ******************************************************************************
 */

package config

import "fmt"

type MysqlConf struct {
	UserName     string
	Password     string
	Host         string
	Port         int
	DBName       string
	Charset      string
	Loc          string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	Enable       bool
}

func (s *MysqlConf) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		s.UserName,
		s.Password,
		s.Host,
		s.DBName,
		s.Charset,
		s.ParseTime,
		s.Loc,
	)
}
