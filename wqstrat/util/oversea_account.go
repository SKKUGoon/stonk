package util

import (
	"encoding/json"
	"errors"
	"fmt"
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
	OutputOne OverseaAccountResponseBodyOutputOne `json:"output1"`
	OutputTwo OverseaAccountResponseBodyOutputTwo `json:"output2"`
}

type OverseaAccountResponseBodyOutputOne struct {
	AccountNumber                 string `json:"cano"`
	AccountProductCode            string `json:"acnt_prdt_cd"`
	OverseaProductNumber          string `json:"ovrs_pdno"`
	OverseaItemName               string `json:"ovrs_item_name"`
	ForeignCurrencyEvaluatedPnl   string `json:"frcr_evlu_pfls_amt"`
	EvaluatedPnl                  string `json:"evlu_pfls_rt"`
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
	TotalEvaluatedPnl              string `json:"tot_evlu_pfls_amt"`
	TotalProfitReturn              string `json:"tot_pftrt"`
	ForeignCurrencyBuyAmountSumOne string `json:"frcr_buy_amt_smtl1"`
	OverseaRealizedPnlTwo          string `json:"ovrs_rlzt_pfls_amt2"`
	ForeignCurrencyBuyAmountSumTwo string `json:"frcr_buy_amt_smtl2"`
}

func (c *KISClient) overseaAccountHeader() OverseaAccountRequestHeader {
	var trId string
	switch c.isTest {
	case false:
		trId = "TTTS3012R"
	case true:
		trId = "VTTS3012R"
	}

	uid := uuid.New()

	// var overseaAccountHeader map[string]string
	header := OverseaAccountRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(),
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         trId,
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaAccountBody() (OverseaAccountRequestQuery, error) {
	result := OverseaAccountRequestQuery{}

	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		return result, errors.New("failed to get account number from environment file")
	}

	result.AccountNumber = acnt[:8]
	result.AccountProductCode = acnt[8:]
	result.OverseaExchange = string(TestNasdaq)             // Test
	result.TransactionCurrency = string(UnitedStatesDollar) // Test

	return result, nil
}

func (c *KISClient) OverseaAccount() (OverseaAccountResponseHeader, OverseaAccountResponseBody, error) {
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

	// Create header for new request.
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
	if query, err := c.overseaAccountBody(); err == nil {
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

	fmt.Println(request.URL.RawQuery)

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
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	if err != nil {
		return resultHeader, resultBody, err
	}

	err = json.Unmarshal(bytes, &resultBody)
	if err != nil {
		return resultHeader, resultBody, err
	}

	return resultHeader, resultBody, nil
}
