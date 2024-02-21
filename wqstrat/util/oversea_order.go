package util

const OverseaOrderUrl string = "/uapi/overseas-stock/v1/trading/order"

type OverseaOrderTxID string

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

type OverseaOrderRequestBody struct {
}
