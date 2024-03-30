package api

import (
	"log"
	"math"
	"net/http"
	"strategy/binance"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type vsBody struct {
	Underlying string `json:"underlying"`
	Callput    string `json:"callput"`
}

type vsElement struct {
	// Minimum ~ Maximum value
	RangeX [2]float64 `json:"lenX"` // X Axis: Strike price
	RangeY [2]float64 `json:"lenY"` // Y Axis: Time (left) to Maturity
	RangeZ [2]float64 `json:"lenZ"` // Z Axis: Implied Vol

	Data []binance.VolatilitySurfaceMapXY `json:"data"`
}

func (b *Business) volatilitySurfaceElement(ctx *gin.Context) {
	body := vsBody{}
	ctx.ShouldBindJSON(&body)

	if strings.ToLower(body.Callput) == "call" {
		callXY, _ := b.Binance.OptionVolatilitySurfaceAxis(binance.OptionCall, body.Underlying)
		x, y, z := sizeUpVSElement(callXY)
		d := mapToArr(callXY)

		ctx.JSON(http.StatusOK, vsElement{
			RangeX: x,
			RangeY: y,
			RangeZ: z,
			Data:   d,
		})
	} else {
		putXY, _ := b.Binance.OptionVolatilitySurfaceAxis(binance.OptionPut, body.Underlying)
		x, y, z := sizeUpVSElement(putXY)
		d := mapToArr(putXY)

		ctx.JSON(http.StatusOK, vsElement{
			RangeX: x,
			RangeY: y,
			RangeZ: z,
			Data:   d,
		})
	}
}

func mapToArr(data map[string]binance.VolatilitySurfaceMapXY) []binance.VolatilitySurfaceMapXY {
	result := []binance.VolatilitySurfaceMapXY{}
	for _, v := range data {
		result = append(result, v)
	}
	return result
}

func sizeUpVSElement(data map[string]binance.VolatilitySurfaceMapXY) ([2]float64, [2]float64, [2]float64) {
	var minX float64 = math.MaxFloat64
	var minY float64 = math.MaxFloat64
	var minZ float64 = math.MaxFloat64

	var maxX float64 = math.MaxFloat64 * -1
	var maxY float64 = math.MaxFloat64 * -1
	var maxZ float64 = math.MaxFloat64 * -1

	for _, v := range data {
		// Axis X - Strike price
		if sp, err := strconv.ParseFloat(v.StrikePrice, 64); err != nil {
			log.Println("error parsing", v.StrikePrice)
		} else {
			if sp < minX {
				minX = sp
			}
			if sp > maxX {
				maxX = sp
			}
		}

		// Axis Y - Time to maturity
		if float64(v.TimeToMature) < minY {
			minY = float64(v.TimeToMature)
		}
		if float64(v.TimeToMature) > maxY {
			maxY = float64(v.TimeToMature)
		}

		// Axis Z - Implied Volaility
		if fv, err := strconv.ParseFloat(v.ImpliedVolatility, 64); err != nil {
			log.Println("error parsing", v.ImpliedVolatility)
		} else {
			if fv < minZ {
				minZ = fv
			}
			if fv > maxZ {
				minZ = fv
			}
		}
	}

	return [2]float64{minX, maxX}, [2]float64{minY, maxY}, [2]float64{minZ, maxZ}
}
