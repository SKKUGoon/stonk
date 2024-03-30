package api

import (
	"strategy/binance"
	"strategy/kis"

	"github.com/gin-gonic/gin"
)

type Business struct {
	Conn *gin.Engine

	// API Clients
	Brokerage *kis.KISClient
	Binance   *binance.BinanceOptionClient
}

func setDevState(state string) {
	switch {
	case state == "deploy":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func Engine(state string) *Business {
	setDevState(state)

	router := gin.Default()
	router.Use(corsMiddleware())

	// KIS
	kc := kis.Default(false)
	kc.UsePrefixFn(kc.SetOAuthSecurityCode)
	kc.UseClosingFn(kc.RemoveOAuthSecuritCode)

	// Binance - Option
	bc := binance.Default(false)

	return &Business{
		Conn:      router,
		Brokerage: kc,
		Binance:   bc,
	}
}

func (b *Business) MountServiceKIS(group *gin.RouterGroup) {
	group.GET("/account/:region", b.accountOversea)
	group.GET("/periodpnl/:region", b.periodProfitOversea)

	group.POST("/order/oversea", b.overseaOrder)
}

func (b *Business) MountServiceBinance(group *gin.RouterGroup) {
	group.POST("/volatility/:underlying", b.volatilitySurfaceElement)
}

func (b *Business) Shutdown() {
	b.Brokerage.Close()
}
