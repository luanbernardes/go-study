package main

import "fmt"

func main() {
	accountBalance := 100

	for {
		choice := selectChoice()

		switch choice {
		case 1:
			fmt.Println("you balance is: ", accountBalance)

		case 2:
			depositAmount := 0

			fmt.Println("How much would you like to deposit? ")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("success! your amount is:", accountBalance)

		case 3:
			withDrawAmount := 0

			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withDrawAmount)
			accountBalance -= withDrawAmount
			fmt.Println("success! your amount is:", accountBalance)

		default:
			fmt.Println("Goodbye!")
			//break just breaks out of the switch statement not the for loop
			return
		}
	}
}

func selectChoice() (choice int) {
	fmt.Println("Please select an option: ")
	fmt.Println("1. check balance")
	fmt.Println("2. deposit money")
	fmt.Println("3. withdraw money")
	fmt.Println("4. exit")

	fmt.Scan(&choice)

	return choice
}
