package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	fm := filemanager.New()
	//cm := cmdmanager.New()
	priceJob, err := prices.NewTaxIncludedPriceJob(fm, taxRates)
	if err != nil {
		fmt.Println("Could not create job")
		fmt.Println(err)
		return
	}

	err = priceJob.Process()

	if err != nil {
		fmt.Println("Could not process job")
		fmt.Println(err)
	}
}
