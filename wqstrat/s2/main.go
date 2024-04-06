package main

import (
	"encoding/json"
	"fmt"
	"strategy/coin/binance"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	client := binance.DefaultBinance(true)

	// Test Connectivity
	client.Connectivity()
	client.ServerTime()

	testOptionContract(client)
}

func testMarkPrice(c *binance.BinanceClient) {
	testAsset := "BTC-240402-70750-C"
	markPrice, err := c.MarkPrice(testAsset)
	fmt.Println("testMarkPrice", markPrice, err)
}

func testOptionContract(c *binance.BinanceClient) {
	contractResult, _ := c.OptionContractsInfo()
	jstr, err := json.Marshal(contractResult)
	fmt.Println("testOptionContract", string(jstr), err)
}

func testOptionRateLimit(c *binance.BinanceClient) {
	rateLimit, _ := c.OptionRateLimits()
	jstr, err := json.Marshal(rateLimit)
	fmt.Println("testOptionRateLimit", string(jstr), err)
}

func testOptionAsset(c *binance.BinanceClient) {
	assetResult, _ := c.OptionAssetsInfo()
	jstr, err := json.Marshal(assetResult)
	fmt.Println("testOptionAsset", string(jstr), err)
}

func testOptionSymbol(c *binance.BinanceClient) {
	const testUnderlying = "BTCUSDT"
	symbolResult, _ := c.OptionSymbolInfo(testUnderlying)
	for _, i := range symbolResult.([]interface{}) {
		fmt.Println("testOptionSymbol", i)
	}
}

func testOptionSymbolMapCall(c *binance.BinanceClient) {
	const testUnderlying = "BTCUSDT"
	symbolResult, _ := c.OptionSymbol(binance.OptionCall, testUnderlying)
	for k := range symbolResult {
		fmt.Println("testOptionSymbol", k)
	}
}

func testOptionSymbolMapPut(c *binance.BinanceClient) {
	const testUnderlying = "BTCUSDT"
	symbolResult, _ := c.OptionSymbol(binance.OptionPut, testUnderlying)
	for k := range symbolResult {
		fmt.Println("testOptionSymbol", k)
	}
}

func testOptionVSXYMap(c *binance.BinanceClient) {
	const testUnderlying = "BTCUSDT"
	axisXY, _ := c.OptionVolatilitySurfaceAxis(binance.OptionCall, testUnderlying)
	for k, v := range axisXY {
		fmt.Println("testOptionVSXYMap", k, v)
	}
}

func testOptionVSXYZ(c *binance.BinanceClient) {
	const testUnderlying = "BTCUSDT"
	axisXY, _ := c.OptionVolatilitySurfaceAxis(binance.OptionCall, testUnderlying)

	symbols := []string{}
	for k := range axisXY {
		symbols = append(symbols, k)
	}

	markprice, err := c.MarkPriceAll(symbols...)
	if err != nil {
		return
	}

	for k, v := range markprice.(map[string]binance.MarkPriceResponseBody) {
		element := axisXY[k]
		element.ImpliedVolatility = v.MarkIV
		fmt.Println("testOption XYZ axis", k, element)
	}
}
