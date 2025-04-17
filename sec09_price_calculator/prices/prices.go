package prices

import (
	"encoding/json"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

const fileName = "prices.txt"

type taxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludePrices map[string]string
}

// NewTaxIncludePricesJob creates a new taxIncludePricesJob instance
func NewTaxIncludePricesJob(taxRate float64) *taxIncludePriceJob {
	return &taxIncludePriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *taxIncludePriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxInlcudePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInlcudePrice)
	}

	job.TaxIncludePrices = result
	job.displayJSON()

	err := filemanager.WriteJSONToFile(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (job *taxIncludePriceJob) LoadData() {
	lines, err := filemanager.ReadLines(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting strings to floats:", err)
		return
	}

	job.InputPrices = prices
}

func (job *taxIncludePriceJob) displayJSON() {
	b, err := json.Marshal(job)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
