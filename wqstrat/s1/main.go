package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strategy/kis"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	client := kis.Default(false)

	defer client.Close()

	// Make first transaction
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseClosingFn(client.RemoveOAuthSecuritCode)

	client.SetDataInfoTx(kis.KISDataRequest{
		StockSymbol: "AAPL", // US0378331005
		ExchangeKey: "us-nasdaq",
	})

	obj, err := client.ExecDataInfo()
	if err != nil {
		log.Println(err)
	}
	bstr, _ := json.Marshal(obj)
	fmt.Println(string(bstr))

	log.Println("data output end")
}
