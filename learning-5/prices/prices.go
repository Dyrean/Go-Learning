package prices

import (
	"fmt"
	"price-calculator/iomanager"
	"price-calculator/utils"
)

type TaxIncludedPriceJob struct {
	IOManager   iomanager.IOManager `json:"-"`
	TaxRates    []float64           `json:"tax_rates"`
	InputPrices []float64           `json:"input_prices"`
	Result      map[string][]string `json:"result"`
}

func (job *TaxIncludedPriceJob) Process() error {
	result := make(map[string][]string)

	for _, taxRate := range job.TaxRates {
		for _, price := range job.InputPrices {
			taxIncludedPrice := price * (1 + taxRate)
			result[fmt.Sprintf("%.2f", price)] = append(result[fmt.Sprintf("%.2f", price)], fmt.Sprintf("%.2f", taxIncludedPrice))
		}
	}

	fmt.Println(result)
	job.Result = result

	err := job.IOManager.Write(job)
	return err
}

func NewTaxIncludedPriceJob(IO iomanager.IOManager, taxRates []float64) (*TaxIncludedPriceJob, error) {
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
		TaxRates:    taxRates,
	}, nil
}
