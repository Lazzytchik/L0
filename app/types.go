package app

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	db "lazzytchik/L0/db/pg"
	"lazzytchik/L0/env"
	"lazzytchik/L0/http"
	"lazzytchik/L0/models"
	"lazzytchik/L0/storage"
	"lazzytchik/L0/stream"
	"log"
	"os"
)

type App struct {
	Server    http.OrderServer
	JetStream stream.JetsStream
	DataBase  db.Postgres
}

func New() *App {
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	logger.SetPrefix("[ APP ]: ")

	loadEnv()

	dataBase := setDataBase(&logger)
	orders, err := dataBase.GetOrders()
	if err != nil {
		logger.Fatalln("Problem with DB:", err)
	}

	memory := storage.New(orders)

	js := stream.JetsStream{}
	js.Connect()

	return &App{
		Server: http.OrderServer{
			Logger:  &logger,
			Storage: memory,
		},
		JetStream: js,
		DataBase:  dataBase,
	}
}

func (app *App) ReceiveSubscriber() {

	app.JetStream.AddStream("order")

	_, err := app.JetStream.SubscribeAsync(func(msg *nats.Msg) {
		app.Server.Logger.Println("Received message:", msg.Data)

		var order models.Order

		json.Unmarshal(msg.Data, &order)

		err := make([]error, 0, 20)
		if order.Validate(err) {
			id, _ := app.DataBase.InsertOrder(order)
			app.Server.Storage.Add(id, order)
		} else {
			app.Server.Logger.Println("Invalid Message:", err)
		}
	})
	if err != nil {
		app.Server.Logger.Println("Subscription error", err)
	}
}

func loadEnv() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Can't load .env file: %s", envErr)
	}
}

func setDataBase(logger *log.Logger) db.Postgres {
	dataBase, err := db.New(env.ExtractDbConfig(), logger)
	if err != nil {
		logger.Fatalln("Problem with DB connection:", err)
	}

	return dataBase
}
