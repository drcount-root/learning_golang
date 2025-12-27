package prices

import (
	// "errors"
	"fmt"

	"example.com/concurrency_and_goroutines/conversion"
	"example.com/concurrency_and_goroutines/iomanager"
)

type taxIncludedPricesInterface map[string]string

type TaxIncludedPriceJob struct {
	// for IOManager - bacause it tells that this field should be excluded from JSON
	// IOManager         filemanager.FileManager    `json:"-"`
	IOManager         iomanager.IOManager        `json:"-"`
	TaxRate           float64                    `json:"tax_rate"`
	InputPrices       []float64                  `json:"input_prices"`
	TaxIncludedPrices taxIncludedPricesInterface `json:"tax_included_prices"`
}

// constructor function
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

// Struct Methods
func (T *TaxIncludedPriceJob) CalculateTaxIncludedPrices(doneChan chan bool, errChannel chan error) {
	err := (*T).LoadData()

	// errChannel <- errors.New("errr") // to simulate err

	if err != nil {
		// return err
		errChannel <- err
		return
	}

	result := make(taxIncludedPricesInterface)
	for _, price := range (*T).InputPrices {
		// Sprintf returns the resulting string.
		amount := price * (1 + (*T).TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", amount)
	}

	(*T).TaxIncludedPrices = result

	fmt.Printf("Tax Rate: %v\n", result)
	(*T).IOManager.WriteJSON(*T)

	// return (*T).IOManager.WriteJSON(*T)

	(*T).IOManager.WriteJSON(*T)

	doneChan <- true
}

func (T *TaxIncludedPriceJob) LoadData() error {
	lines, err := (*T).IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	(*T).InputPrices = prices

	return nil
}
