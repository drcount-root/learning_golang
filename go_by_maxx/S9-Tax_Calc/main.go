package main

import (
	"fmt"

	// "example.com/tax_calc/cmdmanager"
	"example.com/tax_calc/filemanager"
	"example.com/tax_calc/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.CalculateTaxIncludedPrices()

		if err != nil {
			fmt.Println("Could not process")
			fmt.Println(err)
		}

		// priceJob2 := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
	}
}
