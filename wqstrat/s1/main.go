package main

import (
	util "strategy/util"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	client := util.Default(false)
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseSuffixFn(client.RemoveOAuthSecuritCode)

	client.SetTx(client.TxOverseaAccount)
	client.Exec()
	// b, err := client.OAuthWebsocket()
	// if err != nil {
	// 	fmt.Println("wrong")
	// }
	// fmt.Println(b)

	// srv := "korExec"

	// // Start stream
	// client.StartStream(srv)

	// // Start read handler
	// go client.ReadFromSocket(srv)

	// // Start
	// client.Subscribe(srv, "005930")

	// time.Sleep(time.Second * 60)

	// client.Unsubscribe(srv, "005930")

	// client.CloseStream(srv)

	// Account check

	// ac, bc, err := client.OverseaAccount()
	// if err != nil {
	// 	log.Fatalf("error account %v", err)
	// }
	// fmt.Println("account header", ac)
	// fmt.Println("account body", bc)

}
