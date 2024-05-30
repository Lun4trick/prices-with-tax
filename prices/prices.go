package prices
import (
	"strconv"
)

type PriceWithTax struct {
	TaxRate float64 `json:"tax_rate"`
	BasePrices []float64 `json:"base_prices"`
	PricesIncludedTax map[string]float64 `json:"prices_included_tax"`
}

func New(prices []float64, taxRate float64) *PriceWithTax {
	pricesIncludedTax := make(map[string]float64, len(prices))

	for _, price := range prices {
		taxedPrice := price * (1 + (taxRate / 100))
		pricesIncludedTax[strconv.FormatFloat(price, 'f', 2, 64)] = taxedPrice
	}

	return &PriceWithTax{
		TaxRate: taxRate,
		BasePrices: prices,
		PricesIncludedTax: pricesIncludedTax,
	}
}