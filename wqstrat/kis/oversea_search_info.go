package kis

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

const OverseaSearchInfoUrl string = "/uapi/overseas-price/v1/quotations/search-info"

type OverseaSearchInfoRequestQuery struct {
	ProductTypeCode OverseaProductTypeCode `json:"PRDT_TYPE_CD"`
	ProductCode     string                 `json:"PDNO"`
}

type OverseaSearchInfoResponseBody struct {
	OverseaGetResponseBodyBase
	DetailInfo OverseaSearchInfoResponseBodyOutput `json:"output"`
}

type OverseaSearchInfoResponseBodyOutput struct {
	StandardProductCode string `json:"std_pdno"`
	EngName             string `json:"prdt_eng_name"`
	NationCode          string `json:"natn_cd"`
	NationName          string `json:"natn_name"`
	ExchangeMarketCode  string `json:"tr_mket_cd"`
	ExchangeMarketName  string `json:"tr_mket_name"`
	OverseaExchangeCode string `json:"ovrs_excg_cd"`
	OverseaExchangeName string `json:"ovrs_excg_name"`
	CurrencyCode        string `json:"ovrs_papr"`
	CurrencyName        string `json:"crcy_name"`
}

func (c *KISClient) overseaSearchInfoHeader() OverseaRequestHeader {
	uid := uuid.New()

	header := OverseaRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(), // No Bearer?
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         "CTPF1702R",
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaSearchInfoBody(oexc OverseaExchangeCountry, symbol string) OverseaSearchInfoRequestQuery {
	return OverseaSearchInfoRequestQuery{
		ProductTypeCode: oexc.ExchangeInfo.ExchangeNum3ProdCode,
		ProductCode:     symbol,
	}
}

func (c *KISClient) overseaSearchInfo(exchange OverseaExchangeCountry, symbol string) (OverseaResponseHeader, OverseaSearchInfoResponseBody, error) {
	header := c.overseaSearchInfoHeader()
	jstr, _ := json.Marshal(header)
	fmt.Println(string(jstr))

	query := c.overseaSearchInfoBody(exchange, symbol)

	return overseaGETwHB[
		OverseaSearchInfoRequestQuery,
		OverseaSearchInfoResponseBody,
	](
		header,
		query,
		c.isTest,
		OverseaSearchInfoUrl,
	)
}
