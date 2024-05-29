package main

import (
	"fmt"
	"example.com/price_calculator/prices"
)

func main() {
	userInput := prices.ReadUserInput("Enter the prices separated by space:")
	pricesList := prices.ConvertToFloat(userInput)

	fmt.Println(pricesList)

}