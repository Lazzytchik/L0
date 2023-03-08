package main

import (
	"context"
	"lazzytchik/L0/app"
)

func main() {

	testApp := app.New()

	testApp.ReceiveSubscriber()
	err := testApp.Server.Serve(context.Background())
	if err != nil {
		testApp.Server.Logger.Println("Server error:", err)
	}
}
