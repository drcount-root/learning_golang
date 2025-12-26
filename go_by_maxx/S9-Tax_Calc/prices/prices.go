package prices

import (
	"fmt"

	"example.com/tax_calc/conversion"
	"example.com/tax_calc/filemanager"
)

type taxIncludedPricesInterface map[string]string

type TaxIncludedPriceJob struct {
	// for IOManager - bacause it tells that this field should be excluded from JSON
	IOManager         filemanager.FileManager    `json:"-"`
	TaxRate           float64                    `json:"tax_rate"`
	InputPrices       []float64                  `json:"input_prices"`
	TaxIncludedPrices taxIncludedPricesInterface `json:"tax_included_prices"`
}

// constructor function
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

// Struct Methods
func (T *TaxIncludedPriceJob) CalculateTaxIncludedPrices() {
	(*T).LoadData()
	result := make(taxIncludedPricesInterface)
	for _, price := range (*T).InputPrices {
		// Sprintf returns the resulting string.
		amount := price * (1 + (*T).TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", amount)
	}

	(*T).TaxIncludedPrices = result

	fmt.Printf("Tax Rate: %v\n", result)
	(*T).IOManager.WriteJSON(*T)
}

func (T *TaxIncludedPriceJob) LoadData() {
	lines, err := (*T).IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	(*T).InputPrices = prices

}
