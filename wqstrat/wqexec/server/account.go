package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (b *Business) accountKorea() {

}

func (b *Business) accountOversea(ctx *gin.Context) {
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
		b.Brokerage.SetTx(b.Brokerage.TxOverseaAccountJP)
	case "us":
		b.Brokerage.SetTx(b.Brokerage.TxOverseaAccountUS)
	case "cn":
		b.Brokerage.SetTx(b.Brokerage.TxOverseaAccountCN)
	default:
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Region not supported"})
	}
}
