package api

import (
	"log"
	"math"
	"net/http"
	"strategy/coin/binance"
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
	RangeX [2]float64 `json:"rangeX"` // X Axis: Strike price
	RangeY [2]float64 `json:"rangeY"` // Y Axis: Time (left) to Maturity
	RangeZ [2]float64 `json:"rangeZ"` // Z Axis: Implied Vol

	Data []binance.VolatilitySurfaceMapXY `json:"data"`
}

func (b *Business) volatilitySurfaceElement(ctx *gin.Context) {
	body := vsBody{}
	ctx.ShouldBindJSON(&body)

	var cp binance.OptionCallPut = binance.OptionCall

	switch strings.ToLower(body.Callput) {
	case "call":
		cp = binance.OptionCall
	case "put":
		cp = binance.OptionPut
	default:
		log.Println("wrong body", body)
		ctx.JSON(http.StatusBadRequest, nil)
	}

	// Get X, Y Axis - graph construction basis
	xy, _ := b.Binance.OptionVolatilitySurfaceAxis(cp, body.Underlying)

	// Request for mark price + implied volatility
	symbols := mapKeyToArr(xy)
	price, err := b.Binance.MarkPriceAll(symbols...)
	if err != nil {
		log.Println("no price", price, err)
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	// Insert implied volaility to xy
	xyz := map[string]binance.VolatilitySurfaceMapXY{}
	for k, v := range price.(map[string]binance.MarkPriceResponseBody) {
		xyzElement := xy[k]
		xyzElement.ImpliedVolatility = v.MarkIV
		xyz[k] = xyzElement
	}

	x, y, z := sizeUpVSElement(xyz)
	d := mapValueToArr(xyz)

	ctx.JSON(http.StatusOK, vsElement{
		RangeX: x,
		RangeY: y,
		RangeZ: z,
		Data:   d,
	})
}

func mapKeyToArr(data map[string]binance.VolatilitySurfaceMapXY) []string {
	result := []string{}
	for k := range data {
		result = append(result, k)
	}
	return result
}

func mapValueToArr(data map[string]binance.VolatilitySurfaceMapXY) []binance.VolatilitySurfaceMapXY {
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
				maxZ = fv
			}
		}
	}

	return [2]float64{minX, maxX}, [2]float64{minY, maxY}, [2]float64{minZ, maxZ}
}
