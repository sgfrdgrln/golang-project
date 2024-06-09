package main

import (
	"errors"
	"fmt"
	"os"
)

const calculationFile = "calculation.txt"

func storeCalculationToFile(EBT, profit, ratio float64) {

	calculationText := fmt.Sprintf("Earning before tax: $%.2f\nProfit: $%.2f\nRatio: %.2f", EBT, profit, ratio)

	os.WriteFile(calculationFile, []byte(calculationText), 0644)

}

func main() {

	var revenue, expenses, taxRate float64

	// initialize revenue, expenses, and taxRate based on user input

	revenue, expenses, taxRate, err := promptUser(revenue, expenses, taxRate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// calculate the ebt, profit, and ratio

	EBT, profit, ratio := calculate(revenue, expenses, taxRate)

	// format and print the ebt, profit, and ratio

	formatAndPrintText(EBT, profit, ratio)

	storeCalculationToFile(EBT, profit, ratio)

}
func promptUser(revenue, expenses, taxRate float64) (float64, float64, float64, error) {

	for {
		fmt.Print("Enter revenue: ")
		fmt.Scan(&revenue)

		if revenue <= 0 {
			return revenue, expenses, taxRate, errors.New("Invalid revenue!\nClosing program...")
		}

		fmt.Print("Enter expenses: ")
		fmt.Scan(&expenses)

		if expenses <= 0 {
			return revenue, expenses, taxRate, errors.New("Invalid expenses!\nClosing program...")
		}
		fmt.Print("Enter tax rate: ")
		fmt.Scan(&taxRate)
		if taxRate <= 0 {
			return revenue, expenses, taxRate, errors.New("Invalid tax rate!\nClosing program...")
		}
		return revenue, expenses, taxRate, nil
	}
}
func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {

	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)
	ratio := EBT / profit

	return EBT, profit, ratio

}
func formatAndPrintText(EBT, profit, ratio float64) {

	formattedEBT := fmt.Sprintf("Earning before tax: $%.2f\n", EBT)
	formattedProfit := fmt.Sprintf("Profit: $%.2f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)

}
