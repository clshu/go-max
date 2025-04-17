package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChasns := make([]chan error, len(taxRates))
	for i := 0; i < len(taxRates); i++ {
		doneChans[i] = make(chan bool)
		errorChasns[i] = make(chan error)
	}

	for index, taxRate := range taxRates {
		outFilePath := fmt.Sprintf("result_%.0f.json", taxRate*100)
		fm := filemanager.New("prices.txt", outFilePath)
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChasns[index])
		// if err != nil {
		// 	fmt.Println("Could not process the job: ", err)
		// 	return
		// }
	}

	for index := range doneChans {
		select {
		case <-doneChans[index]:
			fmt.Println("Job done!")
		case err := <-errorChasns[index]:
			fmt.Println("Error: ", err)
		}
	}

}
