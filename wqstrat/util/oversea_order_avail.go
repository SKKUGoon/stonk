package util

import (
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

const OverseaCashUrl string = "/uapi/overseas-stock/v1/trading/inquire-psamount"

type OverseaCashRequestBody struct {
	AccountNumber         string `json:"CANO"`
	AccountProductCode    string `json:"ACNT_PRDT_CD"`
	OverseaExchange       string `json:"OVRS_EXCG_CD"`
	OverseaOrderUnitPrice string `json:"OVERS_ORD_UNPR"`
	ItemCode              string `json:"ITEM_CD"`
}

type OverseaCashResponseBody struct {
	ReturnCode  string `json:"rt_cd"`
	MessageCode string `json:"msg_cd"`
	Message     string `json:"msg1"`
}

type OverseaCashResponseBodyOutputOne struct {
	Currency                     string `json:"tr_crcy_cd"`
	OrderPossible                string `json:"ord_psbl_frcr_amt"`
	ReusePossible                string `json:"sll_ruse_psbl_amt"`
	OverseaOrderPossibleAmount   string `json:"ovrs_ord_psbl_amt"`
	MaxOrderPossibleQuantity     string `json:"max_ord_psbl_qty"`
	OrderPossibleAmountExch      string `json:"echm_af_ord_psbl_amt"`
	MaxOrderPossibleQuantityExch string `json:"echm_af_ord_psbl_qty"`
	OrderPossibleQuantity        string `json:"ord_psbl_qty"`
	ExchangeRate                 string `json:"exrt"`
	TotalOrderPossibleAmount     string `json:"frcr_ord_psbl_amt1"`
	TotalOrderPossibleQuantity   string `json:"ovrs_max_ord_psbl_qty"`
}

func (c *KISClient) TxOverseaCashUS() (interface{}, error) {
	_, body, err := c.overseaCash(UnitedStates, 1, "AAPL")
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) overseaCashHeader() OverseaRequestHeader {
	var trId string

	switch c.isTest {
	case false:
		trId = "TTTS3007R"
	case true:
		// Not supported
	default:
	}

	uid := uuid.New()

	header := OverseaRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(),
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         trId,
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaCashBody(exchange OverseaExchange, unitPrice int, unitItem string) OverseaCashRequestBody {
	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		log.Fatalln("failed to get account number from environment file")
	}

	result := OverseaCashRequestBody{
		AccountNumber:         acnt[:8],
		AccountProductCode:    acnt[8:],
		OverseaExchange:       string(exchange),
		OverseaOrderUnitPrice: strconv.Itoa(unitPrice),
		ItemCode:              unitItem, // Famous stock
	}

	return result
}

func (c *KISClient) overseaCash(exchange OverseaExchange, unitPrice int, unitItem string) (OverseaResponseHeader, OverseaCashResponseBody, error) {
	header := c.overseaCashHeader()
	query := c.overseaCashBody(exchange, unitPrice, unitItem)

	resultHeader, resultBody, err := overseaGETwHB[
		OverseaCashRequestBody,
		OverseaCashResponseBody,
	](
		header,
		query,
		c.isTest,
		OverseaCashUrl,
	)

	return resultHeader, resultBody, err
}
