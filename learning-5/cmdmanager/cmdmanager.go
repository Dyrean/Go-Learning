package cmdmanager

import "fmt"

type CMDInputs struct{}

func New() *CMDInputs {
	return &CMDInputs{}
}

func (ci *CMDInputs) Read() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm prices with ENTER, to exit enter 0 price")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (ci *CMDInputs) Write(data any) error {
	fmt.Println(data)
	return nil
}
