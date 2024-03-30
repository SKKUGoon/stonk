package binance

func (c *BinanceOptionClient) OptionSymbolInfo(underlyings ...string) (interface{}, error) {
	info, err := get[OptionInfo](ExchangeInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	if len(underlyings) > 0 {
		// Find necessary elements from the array
		var underlyingMap map[string]bool = map[string]bool{}
		var result []any = []any{}

		for _, u := range underlyings {
			underlyingMap[u] = true
		}

		for _, oc := range info.OptionSymbols {
			if underlyingMap[oc.Underlying] {
				result = append(result, oc)
			}
		}

		return result, nil
	} else {
		// No underlyings entered. Return total array
		return info.OptionSymbols, nil
	}
}

func (c *BinanceOptionClient) OptionSymbol(callPut OptionCallPut, underlyings ...string) (map[string]bool, error) {
	info, err := get[OptionInfo](ExchangeInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	// Find necessary elements from the array
	var underlyingMap map[string]bool = map[string]bool{}
	var result map[string]bool = map[string]bool{}

	if len(underlyings) > 0 {
		for _, u := range underlyings {
			underlyingMap[u] = true
		}
	}

	for _, oc := range info.OptionSymbols {
		if len(underlyings) > 0 {
			if underlyingMap[oc.Underlying] && oc.Side == string(callPut) {
				result[oc.Symbol] = true
			}
		} else {
			if oc.Side == string(callPut) {
				result[oc.Symbol] = true
			}
		}
	}

	return result, nil
}
