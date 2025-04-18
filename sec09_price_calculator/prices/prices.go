package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

// TaxIncludedPriceJob struct represents a job for calculating tax included prices
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOMangager `json:"-"` // skip this field in JSON
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
}

// NewTaxIncludedPriceJob creates a new taxIncludePricesJob instance
func NewTaxIncludedPriceJob(iom iomanager.IOMangager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

// Process calculates the tax included prices and writes the result to a file
func (job *TaxIncludedPriceJob) Process(done chan bool, errorChan chan error) {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxInlcudePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInlcudePrice)
	}

	job.TaxIncludedPrices = result
	// job.displayJSON()

	err = job.IOManager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return
	}

	done <- true
}

// LoadData reads the input prices from a file and converts them to a slice of floats
func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}
