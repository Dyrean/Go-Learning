package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	calcProfit()
	calcInvestmentAmount()
}

func calcInvestmentAmount() {
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}

func calcProfit() {
	revenue, expenses, taxRate := getProfitInputs()

	earningsBeforeTax, profit, ratio := calcEBTProfitAndRatio(revenue, expenses, taxRate)

	printProfit(earningsBeforeTax, profit, ratio)
}

func getProfitInputs() (float64, float64, float64) {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter Renenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)
	return revenue, expenses, taxRate
}

func calcEBTProfitAndRatio(revenue, expenses, taxRate float64) (earningsBeforeTax float64, profit float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax - (earningsBeforeTax * taxRate)
	ratio = earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio
}

func printProfit(earningsBeforeTax, profit, ratio float64) {
	fmt.Printf("Earning Before Tax (EBT): %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio between EBT and Profit: %.2f\n", ratio)
}
