package binance

const MarkPriceUrl string = "/eapi/v1/mark"

type MarkPriceRequestQuery struct {
	Symbol string `json:"symbol"`
}

type MarkPriceResponseBody struct {
	Symbol         string `json:"symbol"`
	MarkPrice      string `json:"markPrice"`      // Mark price
	BidIV          string `json:"bidIV"`          // Implied volatility Buy
	AskIV          string `json:"askIV"`          // Implied volatility Sell
	MarkIV         string `json:"markIV"`         // Implied volatility mark
	Delta          string `json:"delta"`          // delta
	Theta          string `json:"theta"`          // theta
	Gamma          string `json:"gamma"`          // gamma
	Vega           string `json:"vega"`           // vega
	HighPriceLimit string `json:"highPriceLimit"` // Current highest buy price
	LowPriceLimit  string `json:"lowPriceLimit"`  // Current lowest sell price
}

func (c *BinanceClient) MarkPriceAll(symbols ...string) (interface{}, error) {
	if len(symbols) <= 0 {
		return get[[]MarkPriceResponseBody](MarkPriceUrl, nil)
	}

	var symbolMap map[string]bool = map[string]bool{}
	var result map[string]MarkPriceResponseBody = map[string]MarkPriceResponseBody{}

	for _, s := range symbols {
		symbolMap[s] = true
	}

	prices, err := get[[]MarkPriceResponseBody](MarkPriceUrl, nil)
	if err != nil {
		return nil, err
	}

	for _, p := range prices {
		if symbolMap[p.Symbol] {
			result[p.Symbol] = p
		}
	}

	return result, nil
}

func (c *BinanceClient) MarkPrice(symbol string) (interface{}, error) {
	asset := MarkPriceRequestQuery{Symbol: symbol}
	return get[[]MarkPriceResponseBody](MarkPriceUrl, asset)
}
