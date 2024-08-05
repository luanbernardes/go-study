package main

import "fmt"

func main() {
	revenue := 0.0
	expenses := 0.0
	tax := 0.0

	fmt.Println("Revenue: ")
	_, err := fmt.Scan(&revenue)
	if err != nil {
		fmt.Print("error revenue value!")
		return
	}

	fmt.Println("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("tax: ")
	fmt.Scan(&tax)

	ebt := revenue + expenses
	profit := ebt * (1 - tax/100)

	fmt.Println("Ratio:", ebt/profit)
}
