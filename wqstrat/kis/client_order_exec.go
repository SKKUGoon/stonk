package kis

import (
	"fmt"
	"log"
	"strconv"
)

type OverseaOrderBuySell bool

const (
	OverseaBuy  OverseaOrderBuySell = true
	OverseaSell OverseaOrderBuySell = false
)

type KISOverseaOrder struct {
	Stock         string  `json:"stock"`
	Quantity      float64 `json:"qty"`
	Price         float64 `json:"prc"`
	OrderExecType string  `json:"exec"`
	OrderExchange string  `json:"exch"`
	OrderID       string  `json:"id,omitempty"`

	QtyStr        string
	PrcStr        string
	orderExchange OverseaExchangeCountry
	buySell       OverseaOrderBuySell
}

func CreateFxExcOrder(
	stock string,
	qty float64,
	prc float64,
	exch string,
	exec string,
	buysell OverseaOrderBuySell, // Buy if true, Sell if false
) *KISOverseaOrder {
	ex, ok := FxExchangeMap[OverseaFxExKey(exch)]
	if !ok {
		log.Panicln("no exchange found")
	}

	ot, ok := FxOrderType[exec]
	if !ok {
		// Defaulted to "00"
		ot = ""
	}

	qtyStr := strconv.FormatFloat(qty, 'f', ex.precisionPointQty, 64)
	prcStr := strconv.FormatFloat(prc, 'f', ex.precisionPointPrc, 64)

	order := KISOverseaOrder{
		buySell:       buysell,
		Stock:         stock,
		QtyStr:        qtyStr,
		PrcStr:        prcStr,
		OrderExecType: string(ot),

		orderExchange: ex,
		Price:         prc,
		Quantity:      qty,
	}

	return &order
}

func (o *KISOverseaOrder) CreateFxExcOrder() *KISOverseaOrder {
	if _, ok := fxOrderTypePurchase[o.OrderExecType]; ok {
		return CreateFxExcOrder(o.Stock, o.Quantity, o.Price, o.OrderExchange, o.OrderExecType, true)
	} else if _, ok := fxOrderTypeSell[o.OrderExecType]; ok {
		return CreateFxExcOrder(o.Stock, o.Quantity, o.Price, o.OrderExchange, o.OrderExecType, false)
	} else {
		log.Fatalf("%s not appropriate order exec type", o.OrderExecType)
		return nil
	}
}

func (c *KISClient) SetOrderTx(orderSheet KISOverseaOrder) {
	c.overseaOrders = append(c.overseaOrders, orderSheet)
}

func (c *KISClient) ShowOrderBacklog() map[string]interface{} {
	for i, ord := range c.overseaOrders {
		fmt.Println(i, ord)
	}
	return nil
}

func (c *KISClient) ExecOrderOversea() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute all orders
	payload := map[string]interface{}{}
	for i, ord := range c.overseaOrders {
		_, body, err := c.overseaOrder(ord)
		if err != nil {
			fmt.Printf("err during executing order %v(%s): %v\n", i, ord.Stock, err)
			return nil, err
		} else {
			payload[fmt.Sprintf("payload%v", i)] = body
		}
	}

	// Re-initialize handlers
	c.overseaOrders = []KISOverseaOrder{}
	return payload, nil
}
