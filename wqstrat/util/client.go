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

type WebsocketAuth struct {
	AppKey    string `json:"appkey"`
	SecretKey string `json:"secretkey"`
}

type RESTAuth struct {
	AppKey    string `json:"appkey"`
	AppSecret string `json:"appsecret"`
}

type HandlerFunc func()
type HandlerChain []HandlerFunc

type KISClient struct {
	context.Context
	UserInfoWS    WebsocketAuth
	UserInfoREST  RESTAuth
	KeyExpiration time.Time

	// Handlers that are not using any variable
	handlers HandlerChain

	isTest bool

	OAuthKey       string
	OAuthKeyExpire time.Time

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
	expire, err := time.Parse(time.DateOnly, os.Getenv("__KIS_EXPIRE_DATE"))
	if err != nil {
		log.Fatalf("expiration time parse error: %v", err)
		return nil
	}

	client := KISClient{
		Context: ctx,
		UserInfoWS: WebsocketAuth{
			AppKey:    appkey,
			SecretKey: scrkey,
		},
		UserInfoREST: RESTAuth{
			AppKey:    appkey,
			AppSecret: scrkey,
		},
		handlers:      []HandlerFunc{}, // Set empty handler function list
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
	// Check for `.env`'s API KEY's availability
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

func (c *KISClient) checkForKeys() {
	if ok, err := c.isOAuthKeyAvailable(); err != nil && ok {

	}
}

func (c *KISClient) ExecuteTransaction() {
	// Execute all functions inside queue
	for _, f := range c.handlers {
		f()
	}

	// Re-initialize handlers
	c.handlers = []HandlerFunc{}
}

func (c *KISClient) setTransaction() {

}
