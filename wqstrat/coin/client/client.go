package coin

import (
	"context"
	"fmt"
	"log"
)

type RESTAuth struct {
	AppKey    string
	AppSecret string
}

type HandlerFunc func() (any, error)

type CoinClient struct {
	context.Context
	UserInfo RESTAuth

	// Handlers that are not using any variable
	PreHandlers  []HandlerFunc
	Handlers     []HandlerFunc
	PostHandlers []HandlerFunc

	// Orders
	IsTest bool
}

type CoinExchange struct {
	context.Context
	Services map[ExchangeNames]*CoinClient
}

func Default() *CoinExchange {
	return &CoinExchange{
		Context: context.Background(),

		// Initialize empty client map
		Services: map[ExchangeNames]*CoinClient{},
	}
}

func (ex *CoinExchange) UsePrefixFn(fn HandlerFunc, toExch ExchangeNames) {
	if exchPtr, ok := ex.Services[toExch]; ok {
		exchPtr.usePrefixFn(fn)
	} else {
		log.Printf("no exchange name `%s`. Prefix order voided", toExch)
	}
}

func (ex *CoinExchange) SetTx(fn HandlerFunc, toExch ExchangeNames) {
	if exchPtr, ok := ex.Services[toExch]; ok {
		exchPtr.setTx(fn)
	} else {
		log.Printf("no exchange name `%s`. SetTx order voided", toExch)
	}
}

func (ex *CoinExchange) UsePostFn(fn HandlerFunc, toExch ExchangeNames) {
	if exchPtr, ok := ex.Services[toExch]; ok {
		exchPtr.usePostFn(fn)
	} else {
		log.Printf("no exchange name `%s`. Appendix order voided", toExch)
	}
}

func (ex *CoinExchange) Execute(fn HandlerFunc, toExch ExchangeNames) {
	if exchPtr, ok := ex.Services[toExch]; ok {
		exchPtr.exec()
	} else {
		log.Printf("no exchange name `%s`. Failed to execute", toExch)
	}
}

func (ex *CoinExchange) Close(fn HandlerFunc, toExch ExchangeNames) {
	if exchPtr, ok := ex.Services[toExch]; ok {
		exchPtr.close()
	} else {
		log.Printf("no exchange name `%s`. Appendix order voided", toExch)
	}
}

/* Internal Setter functions */

func (c *CoinClient) usePrefixFn(fn HandlerFunc) {
	c.PreHandlers = append(c.PreHandlers, fn)
}

func (c *CoinClient) usePostFn(fn HandlerFunc) {
	c.PostHandlers = append(c.PostHandlers, fn)
}

func (c *CoinClient) setTx(fn HandlerFunc) {
	c.Handlers = append(c.Handlers, fn)
}

func (c *CoinClient) exec() (map[string]interface{}, error) {
	for i, pf := range c.PreHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	payload := map[string]interface{}{}
	for i, f := range c.Handlers {
		if data, err := f(); err != nil {
			fmt.Printf("err during handler %v: %v\n", i, err)
			return nil, err
		} else if data != nil {
			payload[fmt.Sprintf("payload%v", i)] = data
		}
	}

	c.Handlers = []HandlerFunc{}
	return payload, nil
}

func (c *CoinClient) close() {
	for i, sf := range c.PostHandlers {
		_, err := sf()
		if err != nil {
			fmt.Printf("err during suffix handler %v: %v\n", i, err)
			return
		}
	}
}
