package main

import (
	"fmt"
	"os"
	"encoding/json"
	"example.com/prices_with_tax/prices"
)

func main() {
	pricesFromInput := prices.ReadUserInput("Enter the prices separated by space:")
	taxratesFromInput := prices.ReadUserInput("Enter the tax rates:")
	pricesList := prices.ConvertToFloat(pricesFromInput)
	taxRatesList := prices.ConvertToFloat(taxratesFromInput)

	for _, taxRate := range taxRatesList {
		taxStruct := prices.New(pricesList, taxRate)

		jsonData, err := json.MarshalIndent(taxStruct, "", "    ")

		if err != nil {
			panic(err)
		}

		filename := fmt.Sprintf("prices_with_tax_%v.json", taxRate)

		os.WriteFile(filename, jsonData, 0644)
	}



	fmt.Println(pricesList)
}