/**
 ******************************************************************************
 * @file           : conf_postgres.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/1/26
 ******************************************************************************
 */

package config

import "fmt"

type PostgresConf struct {
	User        string
	Password    string
	DBName      string
	Schema      string
	Host        string
	Port        int
	SSLMode     string
	Application string
	Enable      bool
}

func (p *PostgresConf) Dsn() string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s ",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.DBName,
		p.SSLMode)
	if p.Schema != "" {
		dsn += fmt.Sprintf("search_path=%s ", p.Schema)
	}
	if p.Application != "" {
		dsn += fmt.Sprintf("application_name=%s ", p.Application)
	}

	return dsn
}
