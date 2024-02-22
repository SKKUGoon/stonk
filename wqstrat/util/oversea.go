package util

type OverseaNation string
type OverseaExchange string
type OverseaCurrency string
type OverseaExchangeCode string
type OverseaOrderTxID string
type OverseaOrderType string

const (
	All                OverseaNation = "000"
	NationUnitedStates OverseaNation = "840"
	NationChina        OverseaNation = "156"
	NationJapan        OverseaNation = "392"
	NationVietnam      OverseaNation = "704"
	NationHongKong     OverseaNation = "344"
)

const (
	UnitedStatesCode     OverseaExchangeCode = "00"
	NasdaqCode           OverseaExchangeCode = "01"
	NewYorkExchangeCode  OverseaExchangeCode = "02"
	PinkSheetsCode       OverseaExchangeCode = "03"
	OverTheCounterBBCode OverseaExchangeCode = "04"
	AmexCode             OverseaExchangeCode = "05"

	ChinaCode     OverseaExchangeCode = "00"
	ShanghaiBCode OverseaExchangeCode = "01"
	ShenZhenBCode OverseaExchangeCode = "02"
	ShanghaiACode OverseaExchangeCode = "03"
	ShenZhenACode OverseaExchangeCode = "04"

	JapanCode OverseaExchangeCode = "01"

	HanoiCode    OverseaExchangeCode = "01"
	HochiminCode OverseaExchangeCode = "02"

	HongKongCode    OverseaExchangeCode = "01"
	HongKongCNYCode OverseaExchangeCode = "02"
	HongKongUSDCode OverseaExchangeCode = "03"
)

const (
	TestNasdaq               OverseaExchange = "NASD"
	TestNewYorkStockExchange OverseaExchange = "NYSE"
	TestAmericanExchange     OverseaExchange = "AMEX"

	UnitedStates         OverseaExchange = "NASD"
	Nasdaq               OverseaExchange = "NAS"
	NewYorkStockExchange OverseaExchange = "NYSE"
	AmericanExchange     OverseaExchange = "AMEX"

	// Test and main are same
	HongKong OverseaExchange = "SEHK"
	Shanghai OverseaExchange = "SHAA"
	ShenZhen OverseaExchange = "SZAA"
	Tokyo    OverseaExchange = "TKSE"
	Hanoi    OverseaExchange = "HASE"
	Hochimin OverseaExchange = "VNSE"
)

const (
	UnitedStatesDollar OverseaCurrency = "USD"
	HongKongDollar     OverseaCurrency = "HKD"
	ChineseYuan        OverseaCurrency = "CNY"
	JapaneseYen        OverseaCurrency = "JPY"
	VietnameseDong     OverseaCurrency = "VND"
)

const (
	USLimitPurchase           OverseaOrderType = "00"
	USPremarketLimitPurchase  OverseaOrderType = "32"
	USPostmarketLimitPurchase OverseaOrderType = "34"

	USLimitSell            OverseaOrderType = "00"
	USPremarketMarketSell  OverseaOrderType = "31"
	USPremarketLimitSell   OverseaOrderType = "32"
	USPostMarketMarketSell OverseaOrderType = "33"
	USPostMarketLimitSell  OverseaOrderType = "34"
)

const (
	UnitedStatesBuyOrder  OverseaOrderTxID = "TTTT1002U"
	UnitedStatesSellOrder OverseaOrderTxID = "TTTT1006U"

	JapanBuyOrder  OverseaOrderTxID = "TTTS0308U"
	JapanSellOrder OverseaOrderTxID = "TTTS0307U"

	ShanghaiBuyOrder  OverseaOrderTxID = "TTTS0202U"
	ShanghaiSellOrder OverseaOrderTxID = "TTTS1005U"

	HongKongBuyOrder  OverseaOrderTxID = "TTTS1002U"
	HongKongSellOrder OverseaOrderTxID = "TTTS1001U"

	ShenZhenBuyOrder  OverseaOrderTxID = "TTTS0305U"
	ShenZhenSellOrder OverseaOrderTxID = "TTTS0304U"

	VietnamBuyOrder   OverseaOrderTxID = "TTTS0311U"
	VietnameSellOrder OverseaOrderTxID = "TTTS0310U"

	TestUnitedStatesBuyOrder  OverseaOrderTxID = "VTTT1002U"
	TestUnitedStatesSellOrder OverseaOrderTxID = "VTTT1001U"

	TestJapanBuyOrder  OverseaOrderTxID = "VTTS0308U"
	TestJapanSellOrder OverseaOrderTxID = "TTTS0307U"

	TestShanghaiBuyOrder  OverseaOrderTxID = "VTTS0202U"
	TestShanghaiSellOrder OverseaOrderTxID = "VTTS1005U"

	TestHongKongBuyOrder  OverseaOrderTxID = "VTTS1002U"
	TestHongKongSellOrder OverseaOrderTxID = "VTTS1001U"

	TestShenZhenBuyOrder  OverseaOrderTxID = "VTTS0305U"
	TestShenZhenSellOrder OverseaOrderTxID = "VTTS0304U"

	TestVietnamBuyOrder   OverseaOrderTxID = "VTTS0311U"
	TestVietnameSellOrder OverseaOrderTxID = "VTTS0310U"
)

type OverseaExchangeCountry struct {
	NationCode   OverseaNation
	Exchange     OverseaExchange
	ExchangeCode OverseaExchangeCode
	Currency     OverseaCurrency
	Order        struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}
	precisionPointQty int
	precisionPointPrc int
	isTest            bool
}

var FxOrderType = map[string]OverseaOrderType{
	"us-buy-limit":      USLimitPurchase,
	"uspre-buy-limit":   USPremarketLimitPurchase,
	"uspost-buy-limit":  USPostmarketLimitPurchase,
	"us-sell-limit":     USLimitSell,
	"uspre-sell-market": USPremarketMarketSell,
	"uspre-sell-limit":  USPremarketLimitSell,
	"uspost-sell-limit": USPostMarketLimitSell,
}

var fxOrderTypePurchase = map[string]OverseaOrderType{
	"us-buy-limit":     USLimitPurchase,
	"uspre-buy-limit":  USPremarketLimitPurchase,
	"uspost-buy-limit": USPostmarketLimitPurchase,
}

var fxOrderTypeSell = map[string]OverseaOrderType{
	"us-sell-limit":     USLimitSell,
	"uspre-sell-market": USPremarketMarketSell,
	"uspre-sell-limit":  USPremarketLimitSell,
	"uspost-sell-limit": USPostMarketLimitSell,
}

var FxExchMap = map[string]OverseaExchangeCountry{
	"us-us":     UnitedStatesFx,
	"us-nasdaq": NasdaqFx,
	"us-nyse":   NewYorkExchangeFx,
	"us-amex":   AmexFx,
	"jp-jp":     JapanFx,
}

var UnitedStatesFx = OverseaExchangeCountry{
	NationCode:   NationUnitedStates,
	Exchange:     UnitedStates,
	ExchangeCode: UnitedStatesCode,
	Currency:     UnitedStatesDollar,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  UnitedStatesBuyOrder,
		SellOrder: UnitedStatesSellOrder,
	},
	precisionPointQty: 0,
	precisionPointPrc: 2,
	isTest:            false,
}

var NasdaqFx = OverseaExchangeCountry{
	NationCode:   NationUnitedStates,
	Exchange:     Nasdaq,
	ExchangeCode: NasdaqCode,
	Currency:     UnitedStatesDollar,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  UnitedStatesBuyOrder,
		SellOrder: UnitedStatesSellOrder,
	},
	precisionPointQty: 0,
	precisionPointPrc: 2,
	isTest:            false,
}

var NewYorkExchangeFx = OverseaExchangeCountry{
	NationCode:   NationUnitedStates,
	Exchange:     NewYorkStockExchange,
	ExchangeCode: NewYorkExchangeCode,
	Currency:     UnitedStatesDollar,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  UnitedStatesBuyOrder,
		SellOrder: UnitedStatesSellOrder,
	},
	precisionPointQty: 0,
	precisionPointPrc: 2,
	isTest:            false,
}

var AmexFx = OverseaExchangeCountry{
	NationCode:   NationUnitedStates,
	Exchange:     AmericanExchange,
	ExchangeCode: AmexCode,
	Currency:     UnitedStatesDollar,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  UnitedStatesBuyOrder,
		SellOrder: UnitedStatesSellOrder,
	},
	precisionPointQty: 0,
	precisionPointPrc: 2,
	isTest:            false,
}

var JapanFx = OverseaExchangeCountry{
	NationCode:   NationJapan,
	Exchange:     Tokyo,
	ExchangeCode: JapanCode,
	Currency:     JapaneseYen,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  JapanBuyOrder,
		SellOrder: JapanSellOrder,
	},
	isTest: false,
}

var HongKongFx = OverseaExchangeCountry{
	NationCode:   NationHongKong,
	Exchange:     HongKong,
	ExchangeCode: HongKongCode,
	Currency:     HongKongDollar,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  HongKongBuyOrder,
		SellOrder: HongKongSellOrder,
	},
	isTest: false,
}

var ShanghaiFx = OverseaExchangeCountry{
	NationCode:   NationChina,
	Exchange:     Shanghai,
	ExchangeCode: ShanghaiACode,
	Currency:     ChineseYuan,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  ShanghaiBuyOrder,
		SellOrder: ShanghaiSellOrder,
	},
	isTest: false,
}

var ShenZhenFx = OverseaExchangeCountry{
	NationCode:   NationChina,
	Exchange:     ShenZhen,
	ExchangeCode: ShanghaiACode,
	Currency:     ChineseYuan,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  ShenZhenBuyOrder,
		SellOrder: ShenZhenSellOrder,
	},
	isTest: false,
}

var HanoiFx = OverseaExchangeCountry{
	NationCode:   NationVietnam,
	Exchange:     Hanoi,
	ExchangeCode: HanoiCode,
	Currency:     VietnameseDong,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  VietnamBuyOrder,
		SellOrder: VietnameSellOrder,
	},
	isTest: false,
}

var HochiminFx = OverseaExchangeCountry{
	NationCode:   NationVietnam,
	Exchange:     Hochimin,
	ExchangeCode: HochiminCode,
	Currency:     VietnameseDong,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  VietnamBuyOrder,
		SellOrder: VietnameSellOrder,
	},
	isTest: false,
}
