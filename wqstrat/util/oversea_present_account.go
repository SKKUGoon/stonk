package util

import (
	"log"
	"os"

	"github.com/google/uuid"
)

const OverseaPresentAccountUrl string = "/uapi/overseas-stock/v1/trading/inquire-present-balance"

type OverseaPresentAccountRequestQuery struct {
	AccountNumber      string `json:"CANO"`
	AccountProductCode string `json:"ACNT_PRDT_CD"`
	WonOrForeign       string `json:"WCRC_FRCR_DVSN_CD"` // 01 외화 02 원화
	NationCode         string `json:"NATN_CD"`           // 000 전체 840 미국 344 홍콩 156 중국 392 일본 704 베트남

	// Request body NATN_CD
	// * NATN_CD 000:
	//   00 전체
	//
	// * NATN_CD 840:
	//   00 전체 | 01 나스닥(NASD) | 02 뉴욕거래소(NYSE) | 03 미국 (PINK SHEETS) | 04 미국 (OTCBB) | 05 아멕스(AMEX)
	//
	// * NATN_CD 156:
	//   00 전체 01 상해B 02 심천B 03 상해A 04 심천A
	//
	// * NATN_CD 392:
	//   01 일본
	//
	// * NATN_CD 704:
	//   01 하노이 거래 02 호치민 거래
	//
	// * NATN_CD 344:
	//   01 홍콩 02 홍콩CNY 03 홍콩USD
	MarketCode string `json:"TR_MKET_CD"`
	StockType  string `json:"INQR_DVSN_CD"` // 00 전체 01 일반해외주식 02 미니스탁
}

type OverseaPresentAccountResponseBody struct {
	OverseaGetResponseBodyBase

	DetailStocks []OverseaPresentAccountResponseBodyOutputOne `json:"output1"`
	DetailTotal  OverseaPresentAccountResponseBodyOutputThree `json:"output3"`
}

type OverseaPresentAccountResponseBodyOutputOne struct {
	StockCode string `json:"pdno"`
	StockName string `json:"prdt_name"`

	Quantity             string `json:"cblc_qty13"`
	ExecPurchaseQuantity string `json:"thdt_buy_ccld_qty1"`
	ExecSellQuantity     string `json:"thdt_sll_ccld_qty1"`
	ExecSumQuantity      string `json:"ccld_qty_smtl1"`

	PurchaseBalanceFx string `json:"frcr_pchs_amt"`
	EvaluateBalanceFx string `json:"frcr_evlu_amt2"`
	PnlFx             string `json:"evlu_pfls_amt2"`
	PnlRate           string `json:"evlu_pfls_rt1"`

	CurrentPrice         string `json:"ovrs_now_pric1"`
	AveragePurchasePrice string `json:"avg_unpr3"`

	ExchangeRate    string `json:"bass_exrt"`
	Currency        string `json:"buy_crcy_cd"` // USD, HKD, CNY, JPY, VND
	OverseaExchange string `json:"ovrs_excg_cd"`

	//
	RemainPurchaseBalance string `json:"pchs_rmnd_wcrc_amt"`
	ExecPurchaseBalanceFx string `json:"thdt_buy_ccld_frcr_amt"`
	ExecSellBalanceFx     string `json:"thdt_sll_ccld_frcr_amt"`
}

type OverseaPresentAccountResponseBodyOutputThree struct {
	PurchaseAmount string `json:"pchs_amt_smtl"`
	EvaluateAmount string `json:"evlu_amt_smtl"`
	EvaluatePnl    string `json:"evlu_pfls_amt_smtl"`

	Deposit      string `json:"dncl_amt"`
	CMADeposit   string `json:"cma_evlu_amt"`
	TotalDeposit string `json:"tot_dncl_amt"`
	Cash         string `json:"wdrw_psbl_tot_amt"`
	FxEval       string `json:"frcr_evlu_tota"`
	FxAvail      string `json:"frcr_use_psbl_amt"`
}

func (c *KISClient) TxOverseaPresentAccountUS() (interface{}, error) {
	_, body, err := c.overseaPresentAccount(UnitedStatesFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPresentAccountNasdaq() (interface{}, error) {
	_, body, err := c.overseaPresentAccount(NasdaqFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPresentAccountNYSE() (interface{}, error) {
	_, body, err := c.overseaPresentAccount(NewYorkExchangeFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPresentAccountAMEX() (interface{}, error) {
	_, body, err := c.overseaPresentAccount(AmexFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPresentAccountJP() (interface{}, error) {
	_, body, err := c.overseaPresentAccount(JapanFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

/* Oversea Account - Based on execution (Present time) */

func (c *KISClient) overseaPresentAccountHeader() OverseaRequestHeader {
	var trId string

	switch c.isTest {
	case false:
		trId = "CTRP6504R"
	case true:
		trId = "VTRP6504R"
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

func (c *KISClient) overseaPresentAccountBody(natl OverseaNation, exchangeCode OverseaExchangeCode) OverseaPresentAccountRequestQuery {
	result := OverseaPresentAccountRequestQuery{}

	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		log.Fatalln("failed to get account number from environment file")
	}

	// Common element
	result.AccountNumber = acnt[:8]
	result.AccountProductCode = acnt[8:]
	result.WonOrForeign = "01"
	result.StockType = "00"
	result.NationCode = string(natl)
	result.MarketCode = string(exchangeCode)
	return result
}

func (c *KISClient) overseaPresentAccount(oversea OverseaExchangeCountry) (OverseaResponseHeader, OverseaPresentAccountResponseBody, error) {
	header := c.overseaPresentAccountHeader()
	query := c.overseaPresentAccountBody(oversea.NationCode, oversea.ExchangeCode)

	resultHeader, resultBody, err := overseaGETwHB[
		OverseaPresentAccountRequestQuery,
		OverseaPresentAccountResponseBody,
	](
		header,
		query,
		c.isTest,
		OverseaPresentAccountUrl,
	)

	return resultHeader, resultBody, err
}
