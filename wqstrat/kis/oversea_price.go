package kis

import "github.com/google/uuid"

const OverseaQuotePriceUrl = "/uapi/overseas-price/v1/quotations/price"

type OverseaQuotePriceQuery struct {
	Auth     string                 `json:"AUTH"`
	Exchange OverseaExchangeEngCode `json:"EXCD"`
	Symbol   string                 `json:"SYMB"`
}

type OverseaQuotePriceResponseBody struct {
	OverseaGetResponseBodyBase

	PriceDetail []OverseaQuotePriceResponseBodyOutputOne `json:"output"`
}

type OverseaQuotePriceResponseBodyOutputOne struct {
	SymbolCode    string `json:"rsym"`
	Precision     string `json:"zdiv"`
	PrevDayPrice  string `json:"base"` // 전일 종가
	PrevDayVolume string `json:"pvol"` // 전일 거래량
	CurrentPrice  string `json:"last"`
	Orderable     string `json:"ordy"`
}

func (c *KISClient) overseaQuotePriceHeader() OverseaRequestHeader {
	uid := uuid.New()

	header := OverseaRequestHeader{
		RESTAuth:              c.UserInfoREST,
		Authorization:         c.getBearerAuthorization(), // No Bearer?
		ContentType:           "application/json; charset=utf-8",
		TransactionID:         "HHDFS00000300",
		GlobalTransactionUUID: uid.String(),
	}
	return header
}

func (c *KISClient) overseaQuotePriceBody(exchange OverseaExchangeEngCode, symbol string) OverseaQuotePriceQuery {
	return OverseaQuotePriceQuery{
		Auth:     "",
		Exchange: exchange,
		Symbol:   symbol,
	}
}

func (c *KISClient) overseaQuotePrice(exchange OverseaExchangeEngCode, symbol string) (OverseaResponseHeader, OverseaQuotePriceResponseBody, error) {
	header := c.overseaQuotePriceHeader()
	query := c.overseaQuotePriceBody(exchange, symbol)

	return overseaGETwHB[
		OverseaQuotePriceQuery,
		OverseaQuotePriceResponseBody,
	](
		header,
		query,
		c.isTest,
		OverseaQuotePriceUrl,
	)
}
