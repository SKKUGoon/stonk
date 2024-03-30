package binance

type VolatilitySurfaceMapXY struct {
	StrikePrice       string `json:"strike"` // X Axis displays strike prices
	TimeToMature      int    `json:"t2mat"`  // Y Axis illustrates time to maturity. Short to long
	ImpliedVolatility string `json:"iv"`
}

// Single Asset's volatility surface elements
func (c *BinanceOptionClient) OptionVolatilitySurfaceAxis(
	callPut OptionCallPut,
	underlying string,
) (map[string]VolatilitySurfaceMapXY, error) {
	info, err := get[OptionInfo](ExchangeInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	var axisMap map[string]VolatilitySurfaceMapXY = map[string]VolatilitySurfaceMapXY{}
	for _, oc := range info.OptionSymbols {
		if oc.Underlying == underlying && oc.Side == string(callPut) {
			axisMap[oc.Symbol] = VolatilitySurfaceMapXY{
				StrikePrice:  oc.StrikePrice,
				TimeToMature: int(oc.ExpiryDate) - int(info.ServerTime),
			}
		}
	}

	return axisMap, nil
}
