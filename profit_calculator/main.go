package main

import "fmt"

func main() {
	revenue, expenses, tax := askQuestions()

	ebt := calculateEbt(revenue, expenses)
	profit := calculateProfit(ebt, tax)

	fmt.Println("Ratio:", ebt/profit)
}

func askQuestions() (revenue float64, expenses float64, tax float64) {
	revenue = 0.0
	expenses = 0.0
	tax = 0.0

	fmt.Println("Revenue: ")
	_, err := fmt.Scan(&revenue)
	if err != nil {
		fmt.Print("error revenue value!")
		return 0.0, 0.0, 0.0
	}

	fmt.Println("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("tax: ")
	fmt.Scan(&tax)

	return revenue, expenses, tax
}

func calculateEbt(revenue, expenses float64) float64 {
	return revenue + expenses
}

func calculateProfit(ebt, tax float64) float64 {
	return ebt * (1 - tax/100)
}
