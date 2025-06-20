package main

import (
	"fmt"
	"os"
)

// storing data in files
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to go bank.")

	for {
		fmt.Println("Welcome to go bank!")
		fmt.Println("what do you want to do?")
		fmt.Println("1. Check the balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Enter your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("invalid amount! Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter the amount of your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You cant withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Thank you for using go bank. See you next time.")
		}
	}
}
