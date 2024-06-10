package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/sgfrdgrln/golang-project/file-handling"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = file_handling.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		file_handling.WriteFloatToFile(accountBalance, accountBalanceFile)
	}

	var choice int

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		displayOptions()

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			checkBalance(accountBalance)

		case 2:
			depositMoney(accountBalance)

		case 3:
			withdrawMoney(accountBalance)
		case 4:
			fmt.Println("Stopping the program.")
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice.")
			continue
		}

	}

}

func depositMoney(balance float64) {

	balance, err := file_handling.GetFloatFromFile(accountBalanceFile)

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
	}

	var depositedAmount float64

	fmt.Print("Deposit amount: ")
	fmt.Scan(&depositedAmount)

	if depositedAmount <= 0 {
		fmt.Println("You did not even deposit anything!")

	} else {
		fmt.Printf("\nYou deposited: $%.2f\n", depositedAmount)
		balance += depositedAmount
		fmt.Printf("\nYour new balance is: $%.2f\n", balance)
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
		return
	}
}
func withdrawMoney(balance float64) {

	balance, err := file_handling.GetFloatFromFile(accountBalanceFile)

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
	}

	var withdrawAmount float64
	fmt.Print("Amount to withdraw: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > balance {
		fmt.Println("Your current balance is not enough to withdraw that amount!")
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
	} else if withdrawAmount <= 0 {
		fmt.Println("You did not withdraw anything!")
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
		return
	} else {
		balance -= withdrawAmount
		fmt.Printf("\nYour new balance is: $%.2f\n", balance)
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
		return
	}
}
func checkBalance(balance float64) {

	balance, err := file_handling.GetFloatFromFile(accountBalanceFile)

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		file_handling.WriteFloatToFile(balance, accountBalanceFile)
	}
	if balance == 0 {
		fmt.Println("You have zero balance.")
	} else {
		fmt.Println("Your balance is: $", balance)
	}

}
