package main

import (
	"fmt"
	"os"
	"strconv"
)

func readBalanceFromFile() int64 {
	data, _ := os.ReadFile("balance.txt")
	balanceText := string(data)
	balance, _ := strconv.Atoi(balanceText)
	fmt.Println("data is: ", balanceText)
	return int64(balance)
}

func writeBalanceToFile(accountBalance int64) {
	value := fmt.Sprint(accountBalance)

	err := os.WriteFile("balance.txt", []byte(value), 0644)
	if err != nil {
		fmt.Println("error from write file")
	}
}

func main() {
	for {
		accountBalance := readBalanceFromFile()

		choice := selectChoice()

		switch choice {
		case 1:
			fmt.Println("you balance is: ", accountBalance)

		case 2:
			depositAmount := 0

			fmt.Println("How much would you like to deposit? ")

			fmt.Scan(&depositAmount)
			sum := accountBalance + int64(depositAmount)
			writeBalanceToFile(sum)
			fmt.Println("success! your amount is:", accountBalance)

		case 3:
			withDrawAmount := 0

			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withDrawAmount)
			sum := accountBalance - int64(withDrawAmount)
			writeBalanceToFile(sum)
			fmt.Println("success! your amount is:", accountBalance)

		default:
			fmt.Println("Goodbye!")
			//break just breaks out of the switch statement not the for loop
			return
		}
	}
}

func selectChoice() (choice int64) {
	fmt.Println("Please select an option: ")
	fmt.Println("1. check balance")
	fmt.Println("2. deposit money")
	fmt.Println("3. withdraw money")
	fmt.Println("4. exit")

	fmt.Scan(&choice)

	return choice
}
