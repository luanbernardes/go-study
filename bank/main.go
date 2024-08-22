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
			accountBalance = readBalanceFromFile()
			fmt.Println("success! your amount is:", accountBalance)

		case 3:
			withDrawAmount := 0

			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withDrawAmount)
			sum := accountBalance - int64(withDrawAmount)
			writeBalanceToFile(sum)
			accountBalance = readBalanceFromFile()
			fmt.Println("success! your amount is:", accountBalance)

		default:
			fmt.Println("Goodbye!")
			//break just breaks out of the switch statement not the for loop
			return
		}
	}
}
