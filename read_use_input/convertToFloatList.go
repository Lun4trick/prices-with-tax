package getInput

import (
	"strconv"
	"strings"
)

func ConvertToFloat(prices string) []float64 {
	pricesList := strings.Split(prices, " ")
	result := make([]float64, len(pricesList))

	for index, price := range pricesList {
		convertedValue, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err)
		}

		result[index] = convertedValue
	}

	return result
}
