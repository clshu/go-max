package prices

import (
	"encoding/json"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

const fileName = "prices.txt"

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"` // skip this field in JSON
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

// NewTaxIncludedPriceJob creates a new taxIncludePricesJob instance
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxInlcudePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInlcudePrice)
	}

	job.TaxIncludedPrices = result
	job.displayJSON()

	err := job.IOManager.WriteResult(job)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
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

func (job *TaxIncludedPriceJob) displayJSON() {
	b, err := json.Marshal(job)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
