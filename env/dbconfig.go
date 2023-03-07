package env

import (
	pg "lazzytchik/L0/db/pg"
	"os"
)

func ExtractDbConfig() pg.DataBaseConfig {
	return pg.DataBaseConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
}
