package combo

import (
	"fmt"

	"example.com/concurrency_and_goroutines/filemanager"
	"example.com/concurrency_and_goroutines/prices"
)

func Combo() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.CalculateTaxIncludedPrices()

		go priceJob.CalculateTaxIncludedPrices(doneChans[index], errChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process")
		// 	fmt.Println(err)
		// }
	}

	/*
		-----------------------------------------------------------------------
		Why `select` is used here ?
		-----------------------------------------------------------------------

		Each tax rate is processed concurrently in its own goroutine.
		For every goroutine, we have two channels:

		- doneChans[index] → signals successful completion
		- errChans[index]  → reports an error, if one occurs

		The `select` statement allows the main goroutine to wait on
		multiple channels simultaneously and react to whichever one
		becomes ready first.

		Behavior:
		- If an error is received first, the error case executes.
		- If the success signal is received first, the done case executes.
		- Once one case runs, the other channel is ignored for that iteration.

		Important:
		- `select` does NOT wait for all channels.
		- It unblocks as soon as any one case is ready.
		- This prevents blocking indefinitely when only one outcome
		  (success or failure) is expected per goroutine.

		This pattern ensures:
		- Non-blocking coordination
		- Clear separation of success and error signals
		- Safe synchronization across multiple concurrent tasks

		Each loop iteration handles exactly one goroutine's outcome.
		-----------------------------------------------------------------------
	*/

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
