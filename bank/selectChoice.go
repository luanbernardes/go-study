package main

import "fmt"

func selectChoice() (choice int64) {
	fmt.Println("Please select an option: ")
	fmt.Println("1. check balance")
	fmt.Println("2. deposit money")
	fmt.Println("3. withdraw money")
	fmt.Println("4. exit")

	fmt.Scan(&choice)

	return choice
}
