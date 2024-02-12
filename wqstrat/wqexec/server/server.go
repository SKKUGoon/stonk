package api

import (
	"strategy/util"

	"github.com/gin-gonic/gin"
)

type Business struct {
	Conn      *gin.Engine
	Brokerage *util.KISClient
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

	client := util.Default(false)
	client.UsePrefixFn(client.SetOAuthSecurityCode)
	client.UseClosingFn(client.RemoveOAuthSecuritCode)

	return &Business{
		Conn:      router,
		Brokerage: client,
	}
}

func (b *Business) MountService(group *gin.RouterGroup) {
	group.GET("/account/:region", b.accountOversea)
}

func (b *Business) Shutdown() {
	b.Brokerage.Close()
}
