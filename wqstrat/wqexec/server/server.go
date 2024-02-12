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

func New(state string) {
	setDevState(state)

	router := gin.Default()
	router.Use(corsMiddleware())

}
