package db

import "fmt"

type DataBaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func (dbc *DataBaseConfig) GetConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbc.Username,
		dbc.Password,
		dbc.Host,
		dbc.Port,
		dbc.DbName,
	)
}
