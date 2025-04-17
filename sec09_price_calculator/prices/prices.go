package prices

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
)

const fileName = "prices.txt"

type taxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	taxIncludePrices map[string]float64
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

	fmt.Println(result)
}

func (job *taxIncludePriceJob) LoadData() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading file:", scanner.Err())
		return
	}

	var prices = make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			return
		}
		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
}
