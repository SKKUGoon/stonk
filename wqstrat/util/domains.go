package util

import "fmt"

const (
	KoreaInvestREST     = "https://openapi.koreainvestment.com:9443"
	TestKoreaInvestREST = "https://openapivts.koreainvestment.com:29443"
)

// Websocket
const (
	RealTimeExecutedKor       = "ws://ops.koreainvestment.com:21000"
	RealTimeExecutedKorScheme = "ws"
	RealTimeExecutedKorHost   = "ops.koreainvestment.com:21000"

	TestRealTimeExecutedKor       = "ws://ops.koreainvestment.com:31000"
	TestRealTimeExecutedKorScheme = "ws"
	TestRealTimeExecutedKorHost   = "ops.koreainvestment.com:31000"
)

const (
	ExecutedUrl = "/tryitout/H0STCNT0"
)

const (
	ExecutedTxID = "H0STCNT0"
)

func whereToRequest(test bool, url string) string {
	switch test {
	case false:
		return fmt.Sprintf("%s%s", KoreaInvestREST, url)
	default:
		return fmt.Sprintf("%s%s", TestKoreaInvestREST, url)
	}
}
