package util

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

type Auth struct {
	AppKey    string `json:"appkey"`
	SecretKey string `json:"secretkey"`
}

type KISClient struct {
	context.Context
	UserInfo      Auth
	KeyExpiration time.Time

	isTest bool

	Streams map[string]*KISStream
}

type KISStream struct {
	context.Context
	Conn   *websocket.Conn
	Cancel context.CancelFunc
}

func Default(test bool) *KISClient {
	ctx := context.Background()

	// Read environement file
	appkey := os.Getenv("__KIS_APP_KEY")
	scrkey := os.Getenv("__KIS_SECRET_KEY")
	expire, err := time.Parse(time.DateOnly, "2025-02-02")
	if err != nil {
		log.Fatalf("expiration time parse error: %v", err)
		return nil
	}

	client := KISClient{
		Context: ctx,
		UserInfo: Auth{
			AppKey:    appkey,
			SecretKey: scrkey,
		},
		isTest:        test,
		KeyExpiration: expire,
		Streams:       map[string]*KISStream{},
	}
	if ok := client.checkExpiration(); !ok {
		return nil
	}
	return &client
}

func (c *KISClient) checkExpiration() bool {
	today := time.Now()
	left := c.KeyExpiration.Sub(today)

	days := math.Floor(left.Hours() / 24)

	switch {
	case days >= 5 && days < 10:
		color.Red(fmt.Sprintf("Days till key expiration: %v days left", int(days)))
		return true
	case days < 5:
		color.Red(fmt.Sprintf("Days till key expiration: %v days left. Update API Keys", int(days)))
		return false
	default:
		return true
	}
}
