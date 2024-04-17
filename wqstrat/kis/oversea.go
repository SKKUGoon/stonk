package kis

type OverseaNation string

type OverseaExchange string
type OverseaExchangeCode string
type OverseaExchangeEngCode string
type OverseaProductTypeCode string

type OverseaCurrency string
type OverseaOrderTxID string
type OverseaOrderType string

type OverseaFxExKey string

/* Exchange info */

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
	NasdaqEngCode             OverseaExchangeEngCode = "NAS"
	NasdaqDayEngCode          OverseaExchangeEngCode = "BAQ"
	AmexEngCode               OverseaExchangeEngCode = "AMS"
	AmexDayEngCode            OverseaExchangeEngCode = "BAA"
	NewYorkExchangeEngCode    OverseaExchangeEngCode = "NYS"
	NewYorkExchangeDayEngCode OverseaExchangeEngCode = "BAY"

	JapanEngCode OverseaExchangeEngCode = "TSE"

	ShanghaiEngCode      OverseaExchangeEngCode = "SHS"
	ShanghaiIndexEngCode OverseaExchangeEngCode = "SHI"
	ShenZhenEngCode      OverseaExchangeEngCode = "SZS"
	ShenZhenIndexEngCode OverseaExchangeEngCode = "SZI"

	HanoiEngCode    OverseaExchangeEngCode = "HNX"
	HochiminEngCode OverseaExchangeEngCode = "HSX"

	HongKongEngCode OverseaExchangeEngCode = "HKS"
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
	NasdaqProdCode               OverseaProductTypeCode = "512" // 미국 나스닥
	NewYorkStockExchangeProdCode OverseaProductTypeCode = "513" // 미국 뉴욕
	AmexProdCode                 OverseaProductTypeCode = "529" // 미국 아멕스

	JapanProdCode       OverseaProductTypeCode = "515" // 일본
	HongKongProdCode    OverseaProductTypeCode = "501" // 홍콩
	HongKongCnyProdCode OverseaProductTypeCode = "543" // 홍콩CNY
	HongKongUsdProdCode OverseaProductTypeCode = "558" // 홍콩USD

	HanoiProdCode    OverseaProductTypeCode = "507" // 베트남 하노이
	HochiminProdCode OverseaProductTypeCode = "508" // 베트남 호치민

	ShanghaiAProdCode OverseaProductTypeCode = "551" // 중국 상해A
	ShenZhenAProdCode OverseaProductTypeCode = "552" // 중국 심천A
)

/* Currency info */

const (
	UnitedStatesDollar OverseaCurrency = "USD"
	HongKongDollar     OverseaCurrency = "HKD"
	ChineseYuan        OverseaCurrency = "CNY"
	JapaneseYen        OverseaCurrency = "JPY"
	VietnameseDong     OverseaCurrency = "VND"
)

/* Order info */

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

type OverseaExchangeInfo struct {
	NationCode           OverseaNation
	ExchangeNum2Code     OverseaExchangeCode
	ExchangeNum3ProdCode OverseaProductTypeCode

	ExchangeEng3Code OverseaExchangeEngCode
	ExchangeEng4Code OverseaExchange
}

type OverseaExchangeCountry struct {
	// Exchange code info
	ExchangeInfo OverseaExchangeInfo

	// Currency code
	Currency OverseaCurrency

	// Order code
	Order struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}
	precisionPointQty int
	precisionPointPrc int

	// Debug
	isTest bool
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

/* One source of truth */

// Keys for map
const (
	USDUnitedStates         OverseaFxExKey = "us-us"
	USDNasdaq               OverseaFxExKey = "us-nasdaq"
	USDNewYorkStockExchange OverseaFxExKey = "us-nyse"
	USDAmericanExchange     OverseaFxExKey = "us-amex"
	JPYTokyo                OverseaFxExKey = "jp-jp"
	ALLHongKong             OverseaFxExKey = "all-hk"
	CNYShanghaiA            OverseaFxExKey = "cn-shanghai-a"
	CNYShanghaiB            OverseaFxExKey = "cn-shanghai-b"
	CNYShenZhenA            OverseaFxExKey = "cn-shenzhen-a"
	CNYShenZhenB            OverseaFxExKey = "cn-shenzhen-b"
	VNDHochimin             OverseaFxExKey = "vn-hochimin"
	VNDHanoi                OverseaFxExKey = "vn-hanoi"
)

var FxExchangeMap = map[OverseaFxExKey]OverseaExchangeCountry{
	// United States
	USDUnitedStates:         UnitedStatesFx,
	USDNasdaq:               NasdaqFx,
	USDNewYorkStockExchange: NewYorkExchangeFx,
	USDAmericanExchange:     AmexFx,

	// Japan
	JPYTokyo: JapanFx,

	// HongKong
	ALLHongKong: HongKongFx,

	// China
	CNYShanghaiA: ShanghaiAFx,
	CNYShanghaiB: ShanghaiBFx,
	CNYShenZhenA: ShenZhenAFx,
	CNYShenZhenB: ShenZhenBFx,

	// Vietnam
	VNDHanoi:    HanoiFx,
	VNDHochimin: HochiminFx,
}

var UnitedStatesFx = OverseaExchangeCountry{
	// Exchange code info
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:       NationUnitedStates,
		ExchangeNum2Code: UnitedStatesCode,
		ExchangeEng4Code: UnitedStates,
	},

	// Currency code
	Currency: UnitedStatesDollar,

	// Order code
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  UnitedStatesBuyOrder,
		SellOrder: UnitedStatesSellOrder,
	},
	precisionPointQty: 0,
	precisionPointPrc: 2,

	// Debug
	isTest: false,
}

var NasdaqFx = OverseaExchangeCountry{
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationUnitedStates,
		ExchangeNum2Code:     NasdaqCode,
		ExchangeNum3ProdCode: NasdaqProdCode,
		ExchangeEng3Code:     NasdaqDayEngCode,
		ExchangeEng4Code:     Nasdaq,
	},

	Currency: UnitedStatesDollar,
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
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationUnitedStates,
		ExchangeNum2Code:     NewYorkExchangeCode,
		ExchangeNum3ProdCode: NewYorkStockExchangeProdCode,
		ExchangeEng3Code:     NewYorkExchangeEngCode,
		ExchangeEng4Code:     NewYorkStockExchange,
	},

	Currency: UnitedStatesDollar,
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
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationUnitedStates,
		ExchangeNum2Code:     AmexCode,
		ExchangeNum3ProdCode: AmexProdCode,
		ExchangeEng3Code:     AmexEngCode,
		ExchangeEng4Code:     AmericanExchange,
	},

	Currency: UnitedStatesDollar,
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
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationJapan,
		ExchangeNum2Code:     JapanCode,
		ExchangeNum3ProdCode: JapanProdCode,
		ExchangeEng3Code:     JapanEngCode,
		ExchangeEng4Code:     Tokyo,
	},

	Currency: JapaneseYen,
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
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationHongKong,
		ExchangeNum2Code:     HongKongCode,
		ExchangeNum3ProdCode: HongKongProdCode,
		ExchangeEng3Code:     HongKongEngCode,
		ExchangeEng4Code:     HongKong,
	},

	Currency: HongKongDollar,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  HongKongBuyOrder,
		SellOrder: HongKongSellOrder,
	},
	isTest: false,
}

var ShanghaiAFx = OverseaExchangeCountry{
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationChina,
		ExchangeNum2Code:     ShanghaiACode,
		ExchangeNum3ProdCode: ShanghaiAProdCode,
		ExchangeEng3Code:     ShanghaiEngCode,
		ExchangeEng4Code:     Shanghai,
	},

	Currency: ChineseYuan,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  ShanghaiBuyOrder,
		SellOrder: ShanghaiSellOrder,
	},
	isTest: false,
}

var ShanghaiBFx = OverseaExchangeCountry{
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:       NationChina,
		ExchangeNum2Code: ShanghaiBCode,
		ExchangeEng3Code: ShanghaiEngCode,
		ExchangeEng4Code: Shanghai,
	},

	Currency: ChineseYuan,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  ShanghaiBuyOrder,
		SellOrder: ShanghaiSellOrder,
	},
	isTest: false,
}

var ShenZhenAFx = OverseaExchangeCountry{
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationChina,
		ExchangeNum2Code:     ShenZhenACode,
		ExchangeNum3ProdCode: ShenZhenAProdCode,
		ExchangeEng3Code:     ShenZhenEngCode,
		ExchangeEng4Code:     ShenZhen,
	},

	Currency: ChineseYuan,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  ShenZhenBuyOrder,
		SellOrder: ShenZhenSellOrder,
	},
	isTest: false,
}

var ShenZhenBFx = OverseaExchangeCountry{
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:       NationChina,
		ExchangeNum2Code: ShenZhenBCode,
		ExchangeEng3Code: ShenZhenEngCode,
		ExchangeEng4Code: ShenZhen,
	},

	Currency: ChineseYuan,
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
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationVietnam,
		ExchangeNum2Code:     HanoiCode,
		ExchangeNum3ProdCode: HanoiProdCode,
		ExchangeEng3Code:     HanoiEngCode,
		ExchangeEng4Code:     Hanoi,
	},

	Currency: VietnameseDong,
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
	ExchangeInfo: OverseaExchangeInfo{
		NationCode:           NationVietnam,
		ExchangeNum2Code:     HochiminCode,
		ExchangeNum3ProdCode: HochiminProdCode,
		ExchangeEng3Code:     HochiminEngCode,
		ExchangeEng4Code:     Hochimin,
	},

	Currency: VietnameseDong,
	Order: struct {
		BuyOrder  OverseaOrderTxID
		SellOrder OverseaOrderTxID
	}{
		BuyOrder:  VietnamBuyOrder,
		SellOrder: VietnameSellOrder,
	},
	isTest: false,
}
