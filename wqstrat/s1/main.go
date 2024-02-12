package main

import (
	"fmt"
	util "strategy/util"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	client := util.Default(false)
	defer client.Close()

	// Make first transaction
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseClosingFn(client.RemoveOAuthSecuritCode)

	client.SetTx(client.TxOverseaAccountUS)
	client.Exec()

	// Make second transaction
	client.SetTx(client.TxOverseaAccountJP)
	client.SetTx(client.TxOverseaAccountUS)

	fmt.Println("wait 12 seconds first, should execute JP, and US only")
	time.Sleep(time.Second * 12)

	client.Exec()
}
