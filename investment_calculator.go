package main

import (
	"fmt"
)

func main() {

	// fmt.Print("revenue: ")
	revenue := getUserInput("revenue: ")
	expenses := getUserInput("expenses: ")
	taxRate := getUserInput("taxrate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)
	fmt.Printf("ebt: %.1f\nprofit: %v\nratio: %.1f\n ", ebt, profit, ratio)
}

func getUserInput(outputText string) float64 {
	var userInput float64
	fmt.Print(outputText)
	fmt.Scan(&userInput)
	return userInput
}

// (accepted values + type) (return values + types)
func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
