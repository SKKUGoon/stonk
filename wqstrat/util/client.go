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

type HandlerFunc func() (any, error)

type KISClient struct {
	context.Context
	UserInfoWS    WebsocketAuth
	UserInfoREST  RESTAuth
	KeyExpiration time.Time

	// Handlers that are not using any variable
	preHandlers     []HandlerFunc
	handlers        []HandlerFunc
	closingHandlers []HandlerFunc

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

		preHandlers:     []HandlerFunc{},
		closingHandlers: []HandlerFunc{},
		handlers:        []HandlerFunc{}, // Set empty handler function list

		isTest:        test,
		KeyExpiration: expire,
		Streams:       map[string]*KISStream{},
	}
	if ok := client.apiKeyExpirationCheck(); !ok {
		return nil
	}
	return &client
}

func (c *KISClient) apiKeyExpirationCheck() bool {
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

func (c *KISClient) UsePrefixFn(fn HandlerFunc) {
	c.preHandlers = append(c.preHandlers, fn)
}

func (c *KISClient) UseClosingFn(fn HandlerFunc) {
	c.closingHandlers = append(c.closingHandlers, fn)
}

func (c *KISClient) SetTx(fn HandlerFunc) {
	c.handlers = append(c.handlers, fn)
}

func (c *KISClient) Exec() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute all main functions inside queue
	payload := map[string]interface{}{}
	for i, f := range c.handlers {
		if data, err := f(); err != nil {
			fmt.Printf("err during handler %v: %v\n", i, err)
			return nil, err
		} else if data != nil {
			// Some data (any or interface{}) was returned from executing function
			payload[fmt.Sprintf("payload%v", i)] = data
		}
	}

	// Re-initialize handlers
	c.handlers = []HandlerFunc{}
	return payload, nil
}

func (c *KISClient) Close() {
	for i, sf := range c.closingHandlers {
		_, err := sf()
		if err != nil {
			fmt.Printf("err during suffix handler %v: %v\n", i, err)
			return
		}
	}
}
