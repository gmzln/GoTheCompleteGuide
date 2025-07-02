package main

import (
	"controlStructure.com/bank/fileOperations"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	greetUser()
	var accountBalance, err = fileOperations.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
	}

	fmt.Println("Welcome to go bank.")
	fmt.Println("Reach us 24/7: ", randomdata.PhoneNumber())

	for {
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
			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Thank you for using go bank. See you next time.")
		}
	}
}
