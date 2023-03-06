package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Problem: %s", envErr)
	}

	config := ExtractFromEnv()

	conn, err := pgx.Connect(context.Background(), config.GetConnString())

	if err != nil {
		log.Fatal("Problem:", err.Error())
	}

	defer conn.Close(context.Background())
}

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

func ExtractFromEnv() DataBaseConfig {
	return DataBaseConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
}
