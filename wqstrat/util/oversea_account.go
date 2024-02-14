package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const OverseaAccountUrl string = "/uapi/overseas-stock/v1/trading/inquire-balance"

type OverseaExchange string
type OverseaCurrency string

const (
	TestNasdaq               OverseaExchange = "NASD"
	TestNewYorkStockExchange OverseaExchange = "NYSE"
	TestAmericanExchange     OverseaExchange = "AMEX"

	UnitedStates         OverseaExchange = "NASD"
	Nasdaq               OverseaExchange = "NAS"
	NewYorkStockExchange OverseaExchange = "NYSE"
	AmericanExchange     OverseaExchange = "AMEX"

	// Test and main are same
	HongKong OverseaExchange = "SEHK"
	Shanghai OverseaExchange = "SHAA"
	ShenZhen OverseaExchange = "SZAA"
	Tokyo    OverseaExchange = "TKSE"
	Hanoi    OverseaExchange = "HASE"
	Hochimin OverseaExchange = "VNSE"
)

const (
	UnitedStatesDollar OverseaCurrency = "USD"
	HongKongDollar     OverseaCurrency = "HKD"
	ChineseYuan        OverseaCurrency = "CNY"
	JapaneseYen        OverseaCurrency = "JPY"
	VietnameseDong     OverseaCurrency = "VND"
)

type OverseaAccountRequestHeader struct {
	RESTAuth
	ContentType           string `json:"content-type,omitempty"`
	Authorization         string `json:"authorization"`
	PersonalSecurityKey   string `json:"personalseckey,omitempty"` // Essential for corporate client
	TransactionID         string `json:"tr_id"`                    // TTTS3012R for Main VTTS3012R for Test
	TransactionContinued  string `json:"tr_cont,omitempty"`
	CustomerType          string `json:"custtype,omitempty"` // B for corporate client, P for individual
	SeqNo                 string `json:"seq_no,omitempty"`   // Essential for corporate client
	MacAddress            string `json:"mac_address,omitempty"`
	PhoneNumber           string `json:"phone_number,omitempty"`
	IPAddress             string `json:"ip_addr,omitempty"` // Essential for corporate client
	HashKey               string `json:"hashkey,omitempty"`
	GlobalTransactionUUID string `json:"gt_uid,omitempty"`
}

type OverseaAccountRequestQuery struct {
	AccountNumber       string `json:"CANO"`
	AccountProductCode  string `json:"ACNT_PRDT_CD"`
	OverseaExchange     string `json:"OVRS_EXCG_CD"`
	TransactionCurrency string `json:"TR_CRCY_CD"`

	// Not Implemented yet
	ContextAreaFK200 string `json:"CTX_AREA_FK200"`
	ContextAreaNK200 string `json:"CTX_AREA_NK200"`
}

type OverseaAccountResponseHeader struct {
	ContentType             string `json:"content-type"`
	TransactionID           string `json:"tr_id"`
	TransactionIsContinuous string `json:"tr_cont"`
	GlobalTransactionUUID   string `json:"gt_uid"`
}

type OverseaAccountResponseBody struct {
	ReturnCode       string `json:"rt_cd"` // 0 if success
	MessageCode      string `json:"msg_cd"`
	Message          string `json:"msg1"`
	ContextAreaFK200 string `json:"ctx_area_fk200"`
	ContextAreaNK200 string `json:"ctx_area_nk200"`

	// Details of account
	OutputOne []OverseaAccountResponseBodyOutputOne `json:"output1"` // For each stock
	OutputTwo OverseaAccountResponseBodyOutputTwo   `json:"output2"` // For account in total
}

type OverseaAccountResponseBodyOutputOne struct {
	AccountNumber                 string `json:"cano"`
	AccountProductCode            string `json:"acnt_prdt_cd"`
	OverseaProductNumber          string `json:"ovrs_pdno"`
	OverseaItemName               string `json:"ovrs_item_name"`
	ForeignCurrencyEvaluatedPnl   string `json:"frcr_evlu_pfls_amt"`
	EvaluatedPnlRate              string `json:"evlu_pfls_rt"`
	AveragePurchasePrice          string `json:"pchs_avg_pric"`
	OverseaQuantity               string `json:"ovrs_cblc_qty"`
	SellOrderPossibleQuantity     string `json:"ord_psbl_qty"`
	ForeignCurrencyPurchaseAmount string `json:"frcr_pchs_amt1"`
	OverseaStockEvaluatedAmount   string `json:"ovrs_stck_evlu_amt"`
	NowPrice                      string `json:"now_pric2"`
	TransactionCurrencyCode       string `json:"tr_crcy_cd"`
	OverseaExchangeCode           string `json:"ovrs_excg_cd"`
	LoanTypeCode                  string `json:"loan_type_cd"`
	LoanDate                      string `json:"loan_dt"`
	LoanExpireDate                string `json:"expd_dt"`
}

type OverseaAccountResponseBodyOutputTwo struct {
	OverseaRealizedPnlOne          string `json:"ovrs_rlzt_pfls_amt"`
	OverseaTotalPnl                string `json:"ovrs_tot_pfls"`
	RealizedEarningsReturn         string `json:"rlzt_erng_rt"`
	TotalEvaluatedBalance          string `json:"tot_evlu_pfls_amt"`
	TotalProfitReturn              string `json:"tot_pftrt"`
	ForeignCurrencyBuyAmountSumOne string `json:"frcr_buy_amt_smtl1"`
	OverseaRealizedPnlTwo          string `json:"ovrs_rlzt_pfls_amt2"`
	ForeignCurrencyBuyAmountSumTwo string `json:"frcr_buy_amt_smtl2"`
}

type WQAccount struct {
	FxBalance  string             `json:"fxbalance"`
	FxTotalPnl string             `json:"fxpnl"`
	Stocks     []accountStockInfo `json:"stocks"`
}

type accountStockInfo struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	FxPnl    string `json:"fxpnl"`
	PnlRate  string `json:"pnlrate"`
	AvgPrice string `json:"avgprc"`
	Quantity string `json:"qty"`
	NotSold  string `json:"notsold"`
}

func (c *KISClient) TxOverseaAccountUS() (interface{}, error) {
	_, body, err := c.overseaAccount(string(UnitedStates), string(UnitedStatesDollar))
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountJP() (interface{}, error) {
	_, body, err := c.overseaAccount(string(Tokyo), string(JapaneseYen))
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountCN() (interface{}, error) {
	_, body, err := c.overseaAccount(string(Shanghai), string(ChineseYuan))
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func accountInfoTable(data OverseaAccountResponseBody) WQAccount {
	result := WQAccount{
		FxBalance:  data.OutputTwo.TotalEvaluatedBalance,
		FxTotalPnl: data.OutputTwo.OverseaTotalPnl,
		Stocks:     []accountStockInfo{},
	}

	for _, stock := range data.OutputOne {
		cleaned := accountStockInfo{
			Code:     stock.OverseaProductNumber,
			Name:     stock.OverseaItemName,
			FxPnl:    stock.ForeignCurrencyEvaluatedPnl,
			PnlRate:  stock.EvaluatedPnlRate,
			AvgPrice: stock.AveragePurchasePrice,
			Quantity: stock.OverseaQuantity,
			NotSold:  stock.SellOrderPossibleQuantity,
		}

		result.Stocks = append(result.Stocks, cleaned)
	}

	return result
}

/* Korea Investment API Request - Oversea Account */

func (c *KISClient) overseaAccountHeader() OverseaAccountRequestHeader {
	var trId string

	switch c.isTest {
	case false:
		trId = "TTTS3012R"
	case true:
		trId = "VTTS3012R"
	}

	uid := uuid.New()

	header := OverseaAccountRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(),
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         trId,
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaAccountBody(exchange, currency string) (OverseaAccountRequestQuery, error) {
	result := OverseaAccountRequestQuery{}

	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		return result, errors.New("failed to get account number from environment file")
	}

	result.AccountNumber = acnt[:8]
	result.AccountProductCode = acnt[8:]
	result.OverseaExchange = string(exchange)     // Test
	result.TransactionCurrency = string(currency) // Test

	return result, nil
}

func (c *KISClient) overseaAccount(exchange, currency string) (OverseaAccountResponseHeader, OverseaAccountResponseBody, error) {
	var (
		resultHeader OverseaAccountResponseHeader
		resultBody   OverseaAccountResponseBody

		headerMap map[string]string
		queryMap  map[string]string
	)

	request, err := http.NewRequest(
		"GET",
		whereToRequest(c.isTest, OverseaAccountUrl),
		nil,
	)
	if err != nil {
		return resultHeader, resultBody, err
	}

	// Create header for new request
	// Turn struct into map
	if header := c.overseaAccountHeader(); true {
		headerMap, err = structToMap(header)
		if err != nil {
			return resultHeader, resultBody, nil
		}

		for k, v := range headerMap {
			request.Header.Set(k, v)
		}
	}

	// Create body for new request.
	if query, err := c.overseaAccountBody(exchange, currency); err == nil {
		queryMap, err = structToMap(query)
		if err != nil {
			return resultHeader, resultBody, nil
		}

		// Add query elements
		q := request.URL.Query()
		for k, v := range queryMap {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	} else {
		return resultHeader, resultBody, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return resultHeader, resultBody, err
	}
	defer response.Body.Close()

	// Parse the response header
	resultHeader = OverseaAccountResponseHeader{
		TransactionID:           response.Header.Get("tr_id"),
		TransactionIsContinuous: response.Header.Get("tr_cont"),
		GlobalTransactionUUID:   response.Header.Get("gt_uid"),
	}

	// Parse the response body
	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		return resultHeader, resultBody, err
	}

	err = json.Unmarshal(bytes, &resultBody)
	if err != nil {
		return resultHeader, resultBody, err
	}

	return resultHeader, resultBody, nil
}
