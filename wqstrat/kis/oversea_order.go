package kis

import (
	"log"
	"os"
)

const OverseaOrderUrl string = "/uapi/overseas-stock/v1/trading/order"

type OverseaOrderRequestBody struct {
	AccountNumber      string `json:"CANO"`
	AccountProductCode string `json:"ACNT_PRDT_CD"`
	OverseaExchange    string `json:"OVRS_EXCG_CD"`

	StockCode          string `json:"PDNO"`
	OrderQuantity      string `json:"ORD_QTY"`
	OrderPricePerStock string `json:"OVRS_ORD_UNPR"`

	OrderServerCode string `json:"ORD_SVR_DVSN_CD"`
	OrderType       string `json:"ORD_DVSN"`

	// Omitted
	Phone          string `json:"CTAC_TLNO,omitempty"`
	AMCode         string `json:"MGCO_APTM_ODNO,omitempty"`
	PurchaseOrSell string `json:"SLL_TYPE,omitempty"` // if none: purchase, if "00": sell
}

type OverseaOrderResponseBody struct {
	OverseaGetResponseBodyBase
	OrderDetail OverseaOrderResponseBodyOutput
}

type OverseaOrderResponseBodyOutput struct {
	KRXForwardOrderNo string `json:"KRX_FWDG_ORD_ORGNO"`
	OrderNo           string `json:"ODNO"`
	OrderTime         string `json:"ORD_TMD"` // HHMMSS
}

/* Korea Investment API Request - Oversea Account */
func (c *KISClient) overseaOrderHeader(order KISOverseaOrder) OverseaRequestHeader {
	var header OverseaRequestHeader

	if order.buySell == OverseaBuy {
		header = OverseaRequestHeader{
			RESTAuth:      c.UserInfoREST,
			Authorization: c.getBearerAuthorization(),
			ContentType:   "application/json; charset=utf-8",
			TransactionID: string(order.orderExchange.Order.BuyOrder),
		}
	} else {
		header = OverseaRequestHeader{
			RESTAuth:      c.UserInfoREST,
			Authorization: c.getBearerAuthorization(),
			ContentType:   "application/json; charset=utf-8",
			TransactionID: string(order.orderExchange.Order.SellOrder),
		}
	}
	return header
}

func (c *KISClient) overseaOrderBody(order KISOverseaOrder) OverseaOrderRequestBody {
	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		log.Fatalln("failed to get account number from environment file")
	}

	result := OverseaOrderRequestBody{
		AccountNumber:      acnt[:8],
		AccountProductCode: acnt[8:],
		OverseaExchange:    string(order.orderExchange.Exchange),
		StockCode:          order.Stock,
		OrderQuantity:      order.QtyStr,
		OrderPricePerStock: order.PrcStr,
		OrderServerCode:    "0",
		OrderType:          order.OrderExecType,
	}

	return result
}

func (c *KISClient) overseaOrder(order KISOverseaOrder) (OverseaResponseHeader, OverseaOrderResponseBody, error) {
	header := c.overseaOrderHeader(order)
	body := c.overseaOrderBody(order)

	resultHeader, resultBody, err := overseaPOSTwHB[
		OverseaOrderRequestBody,
		OverseaOrderResponseBody,
	](
		header,
		body,
		c.isTest,
		OverseaOrderUrl,
	)

	return resultHeader, resultBody, err
}
