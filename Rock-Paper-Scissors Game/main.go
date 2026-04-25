package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	choices := []string{"rock", "paper", "scissors"}

	userScore := 0
	computerScore := 0

	fmt.Println("Welcome to Rock-Paper-Scissors Game!")

	for {
		userChoice := getUserChoice(choices)
		computerChoice := getComputerChoice(choices)

		winner := getWinner(userChoice, computerChoice)

		switch winner {
		case "tie":
			fmt.Println("It is a tie.")
		case "user":
			fmt.Println("User wins!!!")
			userScore++
		case "computer":
			fmt.Println("Computer wins!!!")
			computerScore++
		}

		fmt.Println("User score:", userScore, "| Computer score:", computerScore)

		if !playAgain() {
			fmt.Println("Thanks for playing...")
			break
		}
	}


}

func getUserChoice(choices []string) string {
	var userChoice string

	for {
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&userChoice)

		userChoice = strings.ToLower(userChoice)

		if isValidChoice(userChoice, choices) {
			return userChoice
		}

		fmt.Println("Invalid choice, please enter rock, paper or scissors.")
	}
}

func isValidChoice(choice string, choices []string) bool {
	for _, c := range choices {
		if choice == c {
			return true
		}
	}
	return false
}

func getComputerChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

func getWinner(user, computer string) string {
	if user == computer {
		return "tie"
	}

	if user == "rock" && computer == "scissors" ||
		user == "paper" && computer == "rock" ||
		user == "scissors" && computer == "paper" {
		return "user"
	}

	return "computer"
}

func playAgain() bool {
	var answer string

	for {
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scanln(&answer)

		answer = strings.ToLower(answer)

		if answer == "y" {
			return true
		} else if answer == "n" {
			return false
		}

		fmt.Println("Please enter 'y' or 'n'")
	}
}



