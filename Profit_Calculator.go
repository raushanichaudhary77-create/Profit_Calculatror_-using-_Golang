package main

import (
	"fmt"
	"strconv"
)

// Function to calculate EBT, Profit and Ratio
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebt := revenue - expenses

	taxAmount := ebt * taxRate / 100
	profit := ebt - taxAmount

	var ratio float64

	if profit != 0 {
		ratio = ebt / profit
	} else {
		ratio = 0
	}

	return ebt, profit, ratio
}

// Function to take input from user
func getUserInput(infoText string) (float64, error) {

	var input string

	fmt.Print(infoText)

	if _, err := fmt.Scanln(&input); err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func main() {

	fmt.Println("===== Profit Calculator =====")

	// Get Revenue
	revenue, err := getUserInput("Enter Revenue: ")

	if err != nil {
		fmt.Println("Invalid revenue:", err)
		return
	}

	// Get Expenses
	expenses, err := getUserInput("Enter Expenses: ")

	if err != nil {
		fmt.Println("Invalid expenses:", err)
		return
	}

	// Get Tax Rate
	taxRate, err := getUserInput("Enter Tax Rate (%): ")

	if err != nil {
		fmt.Println("Invalid tax rate:", err)
		return
	}

	// Calculate financials
	ebt, profit, ratio := calculateFinancials(
		revenue,
		expenses,
		taxRate,
	)

	// Display result
	fmt.Println("\n===== Result =====")

	fmt.Printf("EBT          : %.2f\n", ebt)
	fmt.Printf("Net Profit   : %.2f\n", profit)
	fmt.Printf("Profit Ratio : %.2f\n", ratio)
}
