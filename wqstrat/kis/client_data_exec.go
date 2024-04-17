package kis

import "fmt"

type KISDataRequest struct {
	StockSymbol string `json:"stock"`
	ExchangeKey string `json:"exchange_key"`
}

func (c *KISClient) SetDataQuoteTx(datReq KISDataRequest) {
	_, ok := FxExchangeMap[OverseaFxExKey(datReq.ExchangeKey)]
	if !ok {
		return
	}

	if _, ok = c.overseaQuote[datReq.ExchangeKey]; ok {
		c.overseaQuote[datReq.ExchangeKey] = append(c.overseaQuote[datReq.ExchangeKey], datReq.StockSymbol)
	} else {
		c.overseaQuote[datReq.ExchangeKey] = []string{datReq.StockSymbol}
	}
}

func (c *KISClient) SetDataInfoTx(datReq KISDataRequest) {
	_, ok := FxExchangeMap[OverseaFxExKey(datReq.ExchangeKey)]
	if !ok {
		return
	}

	if _, ok = c.overseaInfo[datReq.ExchangeKey]; ok {
		c.overseaInfo[datReq.ExchangeKey] = append(c.overseaInfo[datReq.ExchangeKey], datReq.StockSymbol)
	} else {
		c.overseaInfo[datReq.ExchangeKey] = []string{datReq.StockSymbol}
	}
}

func (c *KISClient) SetDataDailyTx(datReq KISDataRequest) {
	_, ok := FxExchangeMap[OverseaFxExKey(datReq.ExchangeKey)]
	if !ok {
		return
	}

	if _, ok = c.overseaDaily[datReq.ExchangeKey]; ok {
		c.overseaDaily[datReq.ExchangeKey] = append(c.overseaDaily[datReq.ExchangeKey], datReq.StockSymbol)
	} else {
		c.overseaDaily[datReq.ExchangeKey] = []string{datReq.StockSymbol}
	}
}

func (c *KISClient) SetDataPriceDetailTx(datReq KISDataRequest) {
	_, ok := FxExchangeMap[OverseaFxExKey(datReq.ExchangeKey)]
	if !ok {
		return
	}

	if _, ok = c.overseaPriceDetail[datReq.ExchangeKey]; ok {
		c.overseaPriceDetail[datReq.ExchangeKey] = append(c.overseaPriceDetail[datReq.ExchangeKey], datReq.StockSymbol)
	} else {
		c.overseaPriceDetail[datReq.ExchangeKey] = []string{datReq.StockSymbol}
	}
}

func (c *KISClient) ShowDataQuoteBacklog() map[string][]string {
	return c.overseaQuote
}

func (c *KISClient) ExecDataQuote() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute quote price data request for all symbols
	payload := map[string]interface{}{}
	for exchKey, symbols := range c.overseaQuote {
		exch := FxExchangeMap[OverseaFxExKey(exchKey)]
		for _, s := range symbols {
			_, data, err := c.overseaQuotePrice(exch.ExchangeInfo.ExchangeEng3Code, s)
			if err != nil {
				fmt.Printf("err during data requesting %v", err)
				continue
			}
			payload[s] = data
		}
	}

	// Re-initialize handlers
	c.overseaQuote = map[string][]string{}

	return payload, nil
}

func (c *KISClient) ExecDataInfo() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute quote price data request for all symbols
	payload := map[string]interface{}{}
	for exchKey, symbols := range c.overseaInfo {
		exch := FxExchangeMap[OverseaFxExKey(exchKey)]
		for _, s := range symbols {
			_, data, err := c.overseaSearchInfo(exch, s)
			if err != nil {
				fmt.Printf("err during data requesting %v", err)
				continue
			}
			payload[s] = data
		}
	}

	// Re-initialize handlers
	c.overseaInfo = map[string][]string{}

	return payload, nil
}

func (c *KISClient) ExecDataDaily() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute quote price data request for all symbols
	payload := map[string]interface{}{}
	for exchKey, symbols := range c.overseaDaily {
		exch := FxExchangeMap[OverseaFxExKey(exchKey)]
		for _, s := range symbols {
			// TODO: Change here
			_, data, err := c.overseaQuotePrice(exch.ExchangeInfo.ExchangeEng3Code, s)
			if err != nil {
				fmt.Printf("err during data requesting %v", err)
				continue
			}
			payload[s] = data
		}
	}

	// Re-initialize handlers
	c.overseaDaily = map[string][]string{}

	return payload, nil
}

func (c *KISClient) ExecDataPrcDetail() (map[string]interface{}, error) {
	// Execute all prefix functions - inside the queue
	// Prefix functions are made with `UsePrefixFn`
	for i, pf := range c.preHandlers {
		_, err := pf()
		if err != nil {
			fmt.Printf("err during prefix handler %v: %v\n", i, err)
			return nil, err
		}
	}

	// Execute quote price data request for all symbols
	payload := map[string]interface{}{}
	for exchKey, symbols := range c.overseaPriceDetail {
		exch := FxExchangeMap[OverseaFxExKey(exchKey)]
		for _, s := range symbols {
			// TODO: Change here
			_, data, err := c.overseaQuotePrice(exch.ExchangeInfo.ExchangeEng3Code, s)
			if err != nil {
				fmt.Printf("err during data requesting %v", err)
				continue
			}
			payload[s] = data
		}
	}

	// Re-initialize handlers
	c.overseaPriceDetail = map[string][]string{}

	return payload, nil
}
