package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	var fileName string

	fmt.Println("Enter the file name: ")
	fmt.Scan(&fileName)

	for _, taxRate := range taxRates {
		fm := filemanager.New(fileName, fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cm := cmdmanager.New()
		priceJob, err := prices.NewTaxIncludedPriceJob(fm, taxRate)

		if err != nil {
			fmt.Println("Could not create job")
			fmt.Println(err)
			break
		}

		err = priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
			break
		}
	}
}
