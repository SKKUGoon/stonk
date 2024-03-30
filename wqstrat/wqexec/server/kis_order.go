package api

import (
	"fmt"
	"net/http"
	"strategy/kis"

	"github.com/gin-gonic/gin"
)

func (b *Business) overseaOrder(ctx *gin.Context) {
	orderBody := kis.KISOverseaOrder{}
	defer func() {
		data, err := b.Brokerage.ExecOrderOversea()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to execute brokerage order, function queue not have been cleaned",
			})
			return
		}

		ctx.JSON(http.StatusOK, data)
	}()

	// Parse order request
	ctx.ShouldBindJSON(&orderBody)
	orderConfig := orderBody.CreateFxExcOrder()
	fmt.Println(orderBody, *orderConfig) // Debug purpose. TODO: erase later

	b.Brokerage.SetOrderTx(*orderConfig)
}
