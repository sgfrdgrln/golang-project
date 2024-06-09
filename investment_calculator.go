package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investedAmount float64
	var expectedReturnRate float64
	var years float64

	// initialize based on user input

	investedAmount, expectedReturnRate, years = getUserInput(investedAmount, expectedReturnRate, years)

	// calculate the future value

	futureValue, futureRealValue := calculateFutureValue(investedAmount, expectedReturnRate, years)

	// format then print the future value, and future real value

	formatThenPrint(futureValue, futureRealValue)

}

func calculateFutureValue(investedAmount, expectedReturnRate, years float64) (float64, float64) {

	const inflationRate = 2.5

	futureValue := investedAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue

}

func getUserInput(investedAmount, expectedReturnRate, years float64) (float64, float64, float64) {

	fmt.Print("Enter invested amount: ")
	fmt.Scan(&investedAmount)
	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	return investedAmount, expectedReturnRate, years

}
func formatThenPrint(futureValue, futureRealValue float64) {

	formattedFV := fmt.Sprintf("Future Value: $%.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value (Due to Inflation): $%.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

}
