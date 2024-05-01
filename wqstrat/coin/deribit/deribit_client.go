package deribit

import (
	"log"
	"os"
	coin "strategy/coin/client"
)

type DeribitClient coin.CoinClient

func DefaultDeribit(test bool) *DeribitClient {
	// Read environment file
	scrkey := os.Getenv("__DERIBIT_API_PRIVATE")
	if scrkey == "" {
		log.Panicf("failed to retrieve API information")
	}

	client := DeribitClient{
		UserInfo: coin.RESTAuth{
			AppSecret: scrkey,
		},

		PreHandlers:  []coin.HandlerFunc{},
		Handlers:     []coin.HandlerFunc{},
		PostHandlers: []coin.HandlerFunc{},

		IsTest: test,
	}

	return &client
}
