package util

import (
	"log"
	"os"

	"github.com/google/uuid"
)

const OverseaAccountUrl string = "/uapi/overseas-stock/v1/trading/inquire-balance"

type OverseaAccountRequestQuery struct {
	AccountNumber      string `json:"CANO"`
	AccountProductCode string `json:"ACNT_PRDT_CD"`
	OverseaExchange    string `json:"OVRS_EXCG_CD"`
	Currency           string `json:"TR_CRCY_CD"`

	// Not Implemented yet
	ContextAreaFK200 string `json:"CTX_AREA_FK200"`
	ContextAreaNK200 string `json:"CTX_AREA_NK200"`
}

type OverseaAccountResponseBody struct {
	OverseaGetResponseBodyBase
	ContextAreaFK200 string `json:"ctx_area_fk200"`
	ContextAreaNK200 string `json:"ctx_area_nk200"`

	// Details of account
	DetailStocks  []OverseaAccountResponseBodyOutputOne `json:"output1"` // For each stock
	DetailAccount OverseaAccountResponseBodyOutputTwo   `json:"output2"` // For account in total
}

type OverseaAccountResponseBodyOutputOne struct {
	AccountNumber      string `json:"cano"`
	AccountProductCode string `json:"acnt_prdt_cd"`

	StockCode            string `json:"ovrs_pdno"`
	StockName            string `json:"ovrs_item_name"`
	PnlFx                string `json:"frcr_evlu_pfls_amt"`
	PnlRate              string `json:"evlu_pfls_rt"`
	AveragePurchasePrice string `json:"pchs_avg_pric"`
	Quantity             string `json:"ovrs_cblc_qty"`
	QuantityAvailable    string `json:"ord_psbl_qty"`
	PurchaseBalanceFx    string `json:"frcr_pchs_amt1"`
	EvaluateBalanceFx    string `json:"ovrs_stck_evlu_amt"`
	CurrentPrice         string `json:"now_pric2"`
	Currency             string `json:"tr_crcy_cd"`
	OverseaExchange      string `json:"ovrs_excg_cd"`

	LoanType   string `json:"loan_type_cd"`
	LoanDate   string `json:"loan_dt"`
	LoanExpire string `json:"expd_dt"`
}

type OverseaAccountResponseBodyOutputTwo struct {
	RealizedPnl     string `json:"ovrs_rlzt_pfls_amt"`
	EvaluatePnl     string `json:"ovrs_tot_pfls"`
	RealizedReturn  string `json:"rlzt_erng_rt"`
	EvaluateBalance string `json:"tot_evlu_pfls_amt"`
	PnlRate         string `json:"tot_pftrt"`

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
	_, body, err := c.overseaAccount(UnitedStatesFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountNasdaq() (interface{}, error) {
	_, body, err := c.overseaAccount(NasdaqFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountNYSE() (interface{}, error) {
	_, body, err := c.overseaAccount(NewYorkExchangeFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountAMEX() (interface{}, error) {
	_, body, err := c.overseaAccount(AmexFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountJP() (interface{}, error) {
	_, body, err := c.overseaAccount(JapanFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountCN() (interface{}, error) {
	_, body, err := c.overseaAccount(ShanghaiFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func (c *KISClient) TxOverseaAccountHK() (interface{}, error) {
	_, body, err := c.overseaAccount(HongKongFx)
	if err != nil {
		return body, err
	}

	// Not using header information for now
	return accountInfoTable(body), nil
}

func accountInfoTable(data OverseaAccountResponseBody) WQAccount {
	result := WQAccount{
		FxBalance:  data.DetailAccount.EvaluateBalance,
		FxTotalPnl: data.DetailAccount.EvaluatePnl,
		Stocks:     []accountStockInfo{},
	}

	for _, stock := range data.DetailStocks {
		cleaned := accountStockInfo{
			Code:     stock.StockCode,
			Name:     stock.StockName,
			FxPnl:    stock.PnlFx,
			PnlRate:  stock.PnlRate,
			AvgPrice: stock.AveragePurchasePrice,
			Quantity: stock.Quantity,
			NotSold:  stock.QuantityAvailable,
		}

		result.Stocks = append(result.Stocks, cleaned)
	}

	return result
}

/* Korea Investment API Request - Oversea Account */

func (c *KISClient) overseaAccountHeader() OverseaGetRequestHeader {
	var trId string

	switch c.isTest {
	case false:
		trId = "TTTS3012R"
	case true:
		trId = "VTTS3012R"
	}

	uid := uuid.New()

	header := OverseaGetRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(),
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         trId,
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaAccountBody(exchange OverseaExchange, currency OverseaCurrency) OverseaAccountRequestQuery {
	result := OverseaAccountRequestQuery{}

	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		log.Fatalln("failed to get account number from environment file")
	}

	result.AccountNumber = acnt[:8]
	result.AccountProductCode = acnt[8:]
	result.OverseaExchange = string(exchange)
	result.Currency = string(currency)

	return result
}

func (c *KISClient) overseaAccount(oversea OverseaExchangeCountry) (OverseaGetResponseHeader, OverseaAccountResponseBody, error) {
	header := c.overseaAccountHeader()
	query := c.overseaAccountBody(oversea.Exchange, oversea.Currency)

	resultHeader, resultBody, err := overseaGETwHB[
		OverseaAccountRequestQuery,
		OverseaAccountResponseBody,
	](
		header,
		query,
		c.isTest,
		OverseaAccountUrl,
	)

	return resultHeader, resultBody, err
}
