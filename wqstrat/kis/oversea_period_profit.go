package kis

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const OverseaPeriodProfit string = "/uapi/overseas-stock/v1/trading/inquire-period-profit"

type OverseaPeriodProfitRequestQuery struct {
	AccountNumber      string `json:"CANO"`
	AccountProductCode string `json:"ACNT_PRDT_CD"`

	OverseaExchange OverseaExchange `json:"OVRS_EXCG_CD"`
	NationCode      string          `json:"NATN_CD"`
	CurrencyCode    OverseaCurrency `json:"CRCY_CD"`
	ProductNum      string          `json:"PDNO"`
	StartDate       string          `json:"INQR_STRT_DT"`      // YYYYMMDD
	EndDate         string          `json:"INQR_END_DT"`       // YYYYMMDD
	WonOrForeign    string          `json:"WCRC_FRCR_DVSN_CD"` // 01 외화 02 원화

	// Not Implemented yet
	ContextAreaFK200 string `json:"CTX_AREA_FK200"`
	ContextAreaNK200 string `json:"CTX_AREA_NK200"`
}

type OverseaPeriodProfitResponseBody struct {
	OverseaGetResponseBodyBase

	// Details of profit
	TradeDateTxInfo []OverseaPeriodProfitResponseBodyOutputOne `json:"Output1"`
}

type OverseaPeriodProfitResponseBodyOutputOne struct {
	TradeDay string `json:"trad_day"`

	StockCode            string `json:"ovrs_pdno"`
	StockName            string `json:"ovrs_item_name"`
	ClearedSellQuantity  string `json:"slcl_qty"`
	AveragePurchasePrice string `json:"pchs_avg_pric"`
	PurchaseBalanceFx    string `json:"frcr_pchs_amt1"`
	AverageSellingPrice  string `json:"avg_sll_unpr"` // 평균매도단가
	SellingBalanceFx     string `json:"frcr_sll_amt_smtl1"`
	Fee                  string `json:"stck_sll_tlex"`
	RealizedPnl          string `json:"ovrs_rlzt_pfls_amt"`
	ProfitRate           string `json:"pftrt"`
	OverseaExchange      string `json:"ovrs_excg_cd"`

	ExchangeRate       string `json:"exrt"`
	FirstExchangeRange string `json:"frst_bltn_exrt"`
}

type OverseaPeriodProfitResponseBodyOutputTwo struct {
	SoldAmount     string `json:"stck_sll_amt_smtl"`
	PurchaseAmount string `json:"stck_buy_amt_smtl"`
	StockTradeFee  string `json:"smtl_fee1"`
	RealizedPnl    string `json:"ovrs_rlzt_pfls_tot_amt"`
	PnlRate        string `json:"tot_pftrt"`
	BaseDate       string `json:"bass_dt"`
	ExchangeRate   string `json:"exrt"`
}

func (c *KISClient) TxOverseaPeriodProfitUS() (interface{}, error) {
	_, body, err := c.OverseaPeriodProfit(UnitedStatesFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPeriodProfitNasdaq() (interface{}, error) {
	_, body, err := c.OverseaPeriodProfit(NasdaqFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPeriodProfitNYSE() (interface{}, error) {
	_, body, err := c.OverseaPeriodProfit(NewYorkExchangeFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPeriodProfitJP() (interface{}, error) {
	_, body, err := c.OverseaPeriodProfit(JapanFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPeriodProfitCN() (interface{}, error) {
	_, body, err := c.OverseaPeriodProfit(ShanghaiAFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) TxOverseaPeriodProfitHK() (interface{}, error) {
	_, body, err := c.OverseaPeriodProfit(HongKongFx)
	if err != nil {
		return body, err
	}

	return body, nil
}

/* Korea Investment API Request - Oversea Account Period profit */

func (c *KISClient) overseaPeriodProfitHeader() OverseaRequestHeader {
	var trId string

	// Oversea account's period profit does not offer test
	switch c.isTest {
	case false:
		trId = "TTTS3039R"
	case true:
	}

	uid := uuid.New()

	header := OverseaRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(), // No Bearer?
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         trId,
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaPeriodProfitBody(exchange OverseaExchange, currency OverseaCurrency, pastdays int) OverseaPeriodProfitRequestQuery {

	// Account number
	acnt := os.Getenv("__KIS_ACCOUNT_NUM")
	if acnt == "" {
		log.Fatalln("failed to get account number from environment file")
	}

	now := time.Now().Format("20060102")
	past := time.Now().AddDate(0, 0, -1*pastdays).Format("20060102")

	result := OverseaPeriodProfitRequestQuery{
		AccountNumber:      acnt[:8],
		AccountProductCode: acnt[8:],
		OverseaExchange:    exchange,
		CurrencyCode:       currency,
		WonOrForeign:       "01",
		StartDate:          past,
		EndDate:            now,
	}

	return result
}

func (c *KISClient) OverseaPeriodProfit(oversea OverseaExchangeCountry) (OverseaResponseHeader, OverseaPeriodProfitResponseBody, error) {
	header := c.overseaPeriodProfitHeader()
	query := c.overseaPeriodProfitBody(oversea.ExchangeInfo.ExchangeEng4Code, oversea.Currency, 90)

	resultHeader, resultBody, err := overseaGETwHB[
		OverseaPeriodProfitRequestQuery,
		OverseaPeriodProfitResponseBody,
	](
		header,
		query,
		c.isTest,
		OverseaPeriodProfit,
	)
	return resultHeader, resultBody, err
}
