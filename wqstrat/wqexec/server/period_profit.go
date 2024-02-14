package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (b *Business) periodProfitOversea(ctx *gin.Context) {
	// Client's request parsing
	region := ctx.Param("region")
	defer func() {
		data, err := b.Brokerage.Exec()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to execute brokerage, function queue not have been cleaned",
			})
			return
		}

		ctx.JSON(http.StatusOK, data)
	}()

	switch strings.ToLower(region) {
	case "jp":
		b.Brokerage.SetTx(b.Brokerage.TxOverseaPeriodProfitJP)
	case "us":
		b.Brokerage.SetTx(b.Brokerage.TxOverseaPeriodProfitUS)
	case "cn":
		b.Brokerage.SetTx(b.Brokerage.TxOverseaPeriodProfitCN)
	default:
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Region not supported"})
	}
}
