package main

import (
	"fmt"
)

func main() {

	var revenue, expenses, taxRate float64

	// initialize revenue, expenses, and taxRate based on user input

	revenue, expenses, taxRate = promptUser(revenue, expenses, taxRate)

	// calculate the ebt, profit, and ratio

	EBT, profit, ratio := calculate(revenue, expenses, taxRate)

	// format and print the ebt, profit, and ratio

	formatAndPrintText(EBT, profit, ratio)

}
func promptUser(revenue, expenses, taxRate float64) (float64, float64, float64) {

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	return revenue, expenses, taxRate

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
