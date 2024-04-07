package kis

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

func (c *KISClient) TxOverseaQuotePriceNasdaq(symbol string) (interface{}, error) {
	_, body, err := c.overseaQuotePrice(NasdaqEngCode, symbol)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *KISClient) overseaQuotePriceBody(exchange OverseaExchangeEngCode, symbol string) OverseaQuotePriceQuery {
	return OverseaQuotePriceQuery{
		Auth:     "",
		Exchange: exchange,
		Symbol:   symbol,
	}
}

func (c *KISClient) overseaQuotePrice(exchange OverseaExchangeEngCode, symbol string) (OverseaResponseHeader, OverseaQuotePriceResponseBody, error) {
	header := c.overseaPresentAccountHeader()
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
