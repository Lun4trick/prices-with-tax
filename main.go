package main

import (
	"fmt"
	"example.com/price_calculator/fileops"
)

func main() {
	userInput := fileops.ReadUserInput("Enter the prices separated by space:")
	pricesList := fileops.ConvertToFloat(userInput)

	fmt.Println(pricesList)

}