package binance

import (
	"log"
	"os"
	coin "strategy/coin/client"
)

type BinanceClient coin.CoinClient

func DefaultBinance(test bool) *BinanceClient {
	// Read environment file
	appkey := os.Getenv("__BINANCE_API_PUBLIC")
	scrkey := os.Getenv("__BINANCE_API_PRIVATE")
	if appkey == "" || scrkey == "" {
		log.Panicf("failed to retrieve API information")
	}

	client := BinanceClient{
		UserInfo: coin.RESTAuth{
			AppKey:    appkey,
			AppSecret: scrkey,
		},

		PreHandlers:  []coin.HandlerFunc{},
		Handlers:     []coin.HandlerFunc{},
		PostHandlers: []coin.HandlerFunc{},

		IsTest: test,
	}

	return &client
}
