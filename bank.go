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
		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
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

			break
		case 2:
			depositMoney(accountBalance)

			break
		case 3:
			withdrawMoney(accountBalance)
			break
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
func displayOptions() {

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

}

func depositMoney(accountBalance float64) {

	accountBalance, err := getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		writeBalanceToFile(accountBalance)
	}

	var depositedAmount float64

	fmt.Print("Deposit amount: ")
	fmt.Scan(&depositedAmount)

	if depositedAmount <= 0 {
		fmt.Println("You did not even deposit anything!")

	} else {
		fmt.Printf("\nYou deposited: $%.2f\n", depositedAmount)
		accountBalance += depositedAmount
		fmt.Printf("\nYour new balance is: $%.2f\n", accountBalance)
		writeBalanceToFile(accountBalance)
		return
	}
}
func withdrawMoney(accountBalance float64) {

	accountBalance, err := getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		writeBalanceToFile(accountBalance)
	}

	var withdrawAmount float64
	fmt.Print("Amount to withdraw: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > accountBalance {
		fmt.Println("Your current balance is not enough to withdraw that amount!")
		writeBalanceToFile(accountBalance)
	} else if withdrawAmount <= 0 {
		fmt.Println("You did not withdraw anything!")
		writeBalanceToFile(accountBalance)
		return
	} else {
		accountBalance -= withdrawAmount
		fmt.Printf("\nYour new balance is: $%.2f\n", accountBalance)
		writeBalanceToFile(accountBalance)
		return
	}
}
func checkBalance(accountBalance float64) {

	accountBalance, err := getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		writeBalanceToFile(accountBalance)
	}
	if accountBalance == 0 {
		fmt.Println("You have zero balance.")
	} else {
		fmt.Println("Your balance is: $", accountBalance)
	}

}
