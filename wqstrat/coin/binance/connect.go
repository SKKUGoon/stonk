package binance

const (
	ConnectivityUrl = "/eapi/v1/ping"
	ServerTimeUrl   = "/eapi/v1/time"
	ExchangeInfoUrl = "/eapi/v1/exchangeInfo"
)

type OptionCallPut string

const (
	OptionCall OptionCallPut = "CALL"
	OptionPut  OptionCallPut = "PUT"
)

type OptionInfo struct {
	Timezone   string `json:"timezone,omitempty"`   // Time zone used by the server
	ServerTime int64  `json:"serverTime,omitempty"` // Current system time

	// Option contract underlying asset info
	OptionContracts []struct {
		ID          int    `json:"id,omitempty"`
		BaseAsset   string `json:"baseAsset,omitempty"`   // Base currency
		QuoteAsset  string `json:"quoteAsset,omitempty"`  // Quotation asset
		Underlying  string `json:"underlying,omitempty"`  // Name of the underlying asset of the option contract
		SettleAsset string `json:"settleAsset,omitempty"` // Settlement currency
	} `json:"optionContracts,omitempty"`

	// Option asset info
	OptionAssets []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"` // Asset name
	} `json:"optionAssets,omitempty"`

	// Option trading pair info
	OptionSymbols []struct {
		ContractID int   `json:"contractId,omitempty"`
		ExpiryDate int64 `json:"expiryDate,omitempty"` // expiry time
		Filters    []struct {
			FilterType string `json:"filterType,omitempty"`
			MinPrice   string `json:"minPrice,omitempty"`
			MaxPrice   string `json:"maxPrice,omitempty"`
			TickSize   string `json:"tickSize,omitempty"`
			MinQty     string `json:"minQty,omitempty"`
			MaxQty     string `json:"maxQty,omitempty"`
			StepSize   string `json:"stepSize,omitempty"`
		} `json:"filters,omitempty"`
		ID                   int    `json:"id,omitempty"`
		Symbol               string `json:"symbol,omitempty"`               // Trading pair name
		Side                 string `json:"side,omitempty"`                 // Direction CALL long, PUT short
		StrikePrice          string `json:"strikePrice,omitempty"`          // Strike price
		Underlying           string `json:"underlying,omitempty"`           // Underlying asset of the contract
		Unit                 int    `json:"unit,omitempty"`                 // Contract unit, the quantity of the underlying asset represented by a single contract
		MakerFeeRate         string `json:"makerFeeRate,omitempty"`         // maker commission rate
		TakerFeeRate         string `json:"takerFeeRate,omitempty"`         // taker commission rate
		MinQty               string `json:"minQty,omitempty"`               // Minimum order quantity
		MaxQty               string `json:"maxQty,omitempty"`               // Maximum order quantity
		InitialMargin        string `json:"initialMargin,omitempty"`        // Initial Margin Ratio
		MaintenanceMargin    string `json:"maintenanceMargin,omitempty"`    // Maintenance Margin Ratio
		MinInitialMargin     string `json:"minInitialMargin,omitempty"`     // Min Initial Margin Ratio
		MinMaintenanceMargin string `json:"minMaintenanceMargin,omitempty"` // Min Maintenance Margin Ratio
		PriceScale           int    `json:"priceScale,omitempty"`           // price precision
		QuantityScale        int    `json:"quantityScale,omitempty"`        // quantity precision
		QuoteAsset           string `json:"quoteAsset,omitempty"`           // Quotation asset
	} `json:"optionSymbols,omitempty"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType,omitempty"`
		Interval      string `json:"interval,omitempty"`
		IntervalNum   int    `json:"intervalNum,omitempty"`
		Limit         int    `json:"limit,omitempty"`
	} `json:"rateLimits,omitempty"`
}

// Simple request. No Transactioning - Required
func (c *BinanceClient) Connectivity() (OptionInfo, error) {
	return get[OptionInfo](ConnectivityUrl, nil)
}

func (c *BinanceClient) ServerTime() (OptionInfo, error) {
	return get[OptionInfo](ServerTimeUrl, nil)
}

func (c *BinanceClient) ExchangeInfo() (OptionInfo, error) {
	return get[OptionInfo](ExchangeInfoUrl, nil)
}

/* Derived api calls */

func (c *BinanceClient) OptionContractsInfo(underlyings ...string) (interface{}, error) {
	info, err := get[OptionInfo](ExchangeInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	if len(underlyings) > 0 {
		// Find necessary elements from the array
		var underlyingMap map[string]bool = map[string]bool{} // Turn underlying list into set(map)
		var result []any = []any{}

		for _, u := range underlyings {
			underlyingMap[u] = true
		}

		for _, oc := range info.OptionContracts {
			if underlyingMap[oc.Underlying] {
				result = append(result, oc)
			}
		}

		return result, nil
	} else {
		// No underlyings entered. Return total array
		return info.OptionContracts, nil
	}
}

func (c *BinanceClient) OptionAssetsInfo() (interface{}, error) {
	info, err := get[OptionInfo](ExchangeInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	// Return array
	return info.OptionAssets, nil
}

func (c *BinanceClient) OptionRateLimits() (interface{}, error) {
	info, err := get[OptionInfo](ExchangeInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	return info.RateLimits, nil
}
