package binance

import (
	"context"
	"fmt"
	"os"
)

type RESTAuth struct {
	AppKey    string
	AppSecret string
}

type HandlerFunc func() (any, error)

type BinanceOptionClient struct {
	context.Context
	UserInfo RESTAuth

	// Handlers that are not using any variable
	preHandlers  []HandlerFunc
	handlers     []HandlerFunc
	postHandlers []HandlerFunc

	// Orders

	isTest bool
}

func Default(test bool) *BinanceOptionClient {
	ctx := context.Background()

	// Read environment file
	appkey := os.Getenv("__BINANCE_API_PUBLIC")
	scrkey := os.Getenv("__BINANCE_API_PRIVATE")

	client := BinanceOptionClient{
		Context: ctx,
		UserInfo: RESTAuth{
			AppKey:    appkey,
			AppSecret: scrkey,
		},

		preHandlers:  []HandlerFunc{},
		handlers:     []HandlerFunc{},
		postHandlers: []HandlerFunc{},

		isTest: test,
	}

	return &client
}

func (c *BinanceOptionClient) UsePrefixFn(fn HandlerFunc) {
	c.preHandlers = append(c.preHandlers, fn)
}

func (c *BinanceOptionClient) UsePostFn(fn HandlerFunc) {
	c.postHandlers = append(c.postHandlers, fn)
}

func (c *BinanceOptionClient) SetTx(fn HandlerFunc) {
	c.handlers = append(c.handlers, fn)
}

func (c *BinanceOptionClient) Exec() (map[string]interface{}, error) {
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	payload := map[string]interface{}{}
	for i, f := range c.handlers {
		if data, err := f(); err != nil {
			fmt.Printf("err during handler %v: %v\n", i, err)
			return nil, err
		} else if data != nil {
			payload[fmt.Sprintf("payload%v", i)] = data
		}
	}

	c.handlers = []HandlerFunc{}
	return payload, nil
}

func (c *BinanceOptionClient) Close() {
	for i, sf := range c.postHandlers {
		_, err := sf()
		if err != nil {
			fmt.Printf("err during suffix handler %v: %v\n", i, err)
			return
		}
	}
}
