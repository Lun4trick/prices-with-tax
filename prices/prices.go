package prices

type PriceWithTax struct {
	taxRate float64
	basePrices []float64
	pricesIncludedTax map[float64]float64
}

func New(prices []float64, taxRate float64) *PriceWithTax {
	pricesIncludedTax := make(map[float64]float64, len(prices))

	for _, price := range prices {
		pricesIncludedTax[price] = price * (1 + taxRate)
	}

	return &PriceWithTax{
		taxRate: taxRate,
		basePrices: prices,
		pricesIncludedTax: pricesIncludedTax,
	}
}