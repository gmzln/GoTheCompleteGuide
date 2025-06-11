package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("revenue: ")
	outputText("revenue: ")
	fmt.Scan(&revenue)
	outputText("expenses: ")
	fmt.Scan(&expenses)
	outputText("taxrate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)
	fmt.Printf("ebt: %.1f\nprofit: %v\nratio: %.1f\n ", ebt, profit, ratio)
}

func outputText(text string) {
	fmt.Print(text)
}

// (values + type) (return types)
func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
