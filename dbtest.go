package main

import (
	"context"
	"github.com/joho/godotenv"
	db "lazzytchik/L0/db/pg"
	"lazzytchik/L0/env"
	"lazzytchik/L0/generators"
	"log"
	"os"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Problem: %s", envErr)
	}

	logger := log.Logger{}
	logger.SetPrefix("[ DB ]: ")
	logger.SetOutput(os.Stdout)

	pg, err := db.New(env.ExtractDbConfig(), &logger)

	if err != nil {
		log.Fatal("Problem:", err.Error())
	}

	defer pg.Conn.Close(context.Background())

	order := generators.Order{}.Generate(8)

	id, _ := pg.InsertOrder(order)

	newOrder, _ := pg.GetOrderById(id)

	pg.Logger.Println(newOrder)

}
