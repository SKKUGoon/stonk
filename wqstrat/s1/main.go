package main

import (
	"fmt"
	"log"
	util "strategy/util"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	client := util.Default(false)
	defer client.Close()

	// Make first transaction
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseClosingFn(client.RemoveOAuthSecuritCode)

	client.SetTx(client.TxOverseaPeriodProfitUS)
	client.Exec()

	// // Make second transaction
	// client.SetTx(client.TxOverseaAccountJP)
	// client.SetTx(client.TxOverseaAccountUS)

	// fmt.Println("wait 12 seconds first, should execute JP, and US only")
	// time.Sleep(time.Second * 12)

	data, err := client.Exec()
	if err != nil {
		log.Fatalf("failed to execute client function queue. Queue is not cleaned: %v", err)
	}
	fmt.Println(data)
}
