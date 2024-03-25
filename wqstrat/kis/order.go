package kis

import (
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
	ex, ok := FxExchMap[exch]
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
