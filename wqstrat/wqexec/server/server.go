package api

import (
	"strategy/kis"

	"github.com/gin-gonic/gin"
)

type Business struct {
	Conn      *gin.Engine
	Brokerage *kis.KISClient
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

	client := kis.Default(false)
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseClosingFn(client.RemoveOAuthSecuritCode)

	return &Business{
		Conn:      router,
		Brokerage: client,
	}
}

func (b *Business) MountService(group *gin.RouterGroup) {
	group.GET("/account/:region", b.accountOversea)
	group.GET("/periodpnl/:region", b.periodProfitOversea)

	group.POST("/order/oversea", b.overseaOrder)
}

func (b *Business) Shutdown() {
	b.Brokerage.Close()
}
