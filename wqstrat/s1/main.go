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
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseClosingFn(client.RemoveOAuthSecuritCode)

	client.SetTx(client.TxOverseaAccountUS)
	client.Exec()

	client.SetTx(client.TxOverseaAccountJP)
	client.SetTx(client.TxOverseaAccountUS)

	fmt.Println("wait 12 seconds first, should execute JP, and US only")
	time.Sleep(time.Second * 12)

	client.Exec()

	client.Close()
}
