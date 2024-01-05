package prices

import (
	"fmt"
	"price-calculator/iomanager"
	"price-calculator/utils"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() error {
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
	job.TaxIncludedPrices = result

	err := job.IOManager.Write(job)
	return err
}

func NewTaxIncludedPriceJob(IO iomanager.IOManager, taxRate float64) (*TaxIncludedPriceJob, error) {
	inputLines, err := IO.Read()

	if err != nil {
		fmt.Printf("Failed to read from input\n")
		return nil, err
	}

	inputPrices, err := utils.StringsToFloats(inputLines)

	if err != nil {
		fmt.Printf("Failed to parse string inputs to float")
		return nil, err
	}

	return &TaxIncludedPriceJob{
		IOManager:   IO,
		InputPrices: inputPrices,
		TaxRate:     taxRate,
	}, nil
}
