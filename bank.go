package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		writeBalanceToFile(accountBalance)
	}

	var choice int

	for {
		fmt.Println("Welcome to Go Bank!")

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

	balance, err := getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		writeBalanceToFile(balance)
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
		writeBalanceToFile(balance)
		return
	}
}
func withdrawMoney(balance float64) {

	balance, err := getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		writeBalanceToFile(balance)
	}

	var withdrawAmount float64
	fmt.Print("Amount to withdraw: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > balance {
		fmt.Println("Your current balance is not enough to withdraw that amount!")
		writeBalanceToFile(balance)
	} else if withdrawAmount <= 0 {
		fmt.Println("You did not withdraw anything!")
		writeBalanceToFile(balance)
		return
	} else {
		balance -= withdrawAmount
		fmt.Printf("\nYour new balance is: $%.2f\n", balance)
		writeBalanceToFile(balance)
		return
	}
}
func checkBalance(balance float64) {

	balance, err := getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		writeBalanceToFile(balance)
	}
	if balance == 0 {
		fmt.Println("You have zero balance.")
	} else {
		fmt.Println("Your balance is: $", balance)
	}

}
