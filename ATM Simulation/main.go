package main

import (
	"fmt"
)

//withdraw money
func withdraw(balance int)int{
	var amount int 
	fmt.Println("Enter amount to withdraw: ")
	fmt.Scanln(&amount)

	if amount <= 0{
		fmt.Println("Invalid amount entered!")
	}

	if amount > balance{
		fmt.Println("Insufficient funds!")
		return balance
	}

	balance-=amount
	fmt.Println("Please take your cash...")
	return balance
}
//deposit money
func deposit(balance int)int{
	var amount int
	fmt.Println("Enter amount to deposit: ")
	fmt.Scanln(&amount)
	
	if amount <= 0  {
		fmt.Println("Invalid amount entered!")
		return balance
	}

	balance+=amount
	fmt.Println("Deposit successful.")
	return balance
}
//show balance
func showBalance(balance int){
	fmt.Println("Your current balance is: ",balance)
}

func main() {
	balance := 100000

	for{

		fmt.Println("=============================")
		fmt.Println("        ATM MACHINE")
		fmt.Println("=============================")
		fmt.Println("Current Balance:", balance)
		fmt.Println("=============================")
		fmt.Println("1. Withdraw Cash")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Println("=============================")

		fmt.Println("Select an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice{
		case 1:
			balance = withdraw(balance)
			showBalance(balance)
		case 2:
			balance = deposit(balance)
			showBalance(balance)
		case 3:
			showBalance(balance)
		case 4: 
			fmt.Println("Thank you for using our ATM. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}



