package main

import (
	"fmt"
	util "strategy/util"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.test")

	client := util.Default(true)

	b, err := client.OAuthWebsocket()
	if err != nil {
		fmt.Println("wrong")
	}
	fmt.Println(b)

	srv := "korExec"
	// Start stream
	client.StartStream(srv)

	// Start read handler
	go client.ReadFromSocket(srv)

	// Start
	client.Subscribe(srv, "005930")

	time.Sleep(time.Second * 60)

	client.Unsubscribe(srv, "005930")

	client.CloseStream(srv)
}
