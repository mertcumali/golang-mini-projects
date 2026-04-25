package main

import (
	"fmt"
)

func main(){

	word := "golang"
	hidden := []rune ("______")
	lives := 6
	var wrongGuesses []string

	for lives > 0{

		//clear screen
		fmt.Print("\033[H\033[2J")

		fmt.Println("======================")
		fmt.Println("      HANGMAN")
		fmt.Println("======================")

		drawHangman(lives)

		fmt.Println("Word:", string(hidden))
		fmt.Println("Lives:", lives)
		fmt.Println("Wrong guesses:", wrongGuesses)

		var guess string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&guess)

		found := false

		for i, letter := range word {
			if string(letter) == guess{
				hidden[i] = letter
				found = true
			}
		}

		if !found{
			lives--
			wrongGuesses = append(wrongGuesses, guess)
			fmt.Println("Wrong guess!")
		}

		if string(hidden) == word{
			fmt.Println("🎉 YOU WON! Word:", word)
			return
		}
	}

	fmt.Println("💀 GAME OVER! Word was:", word)

}

func drawHangman(lives int) {
	switch lives {
	case 6:
		fmt.Println(`
  +---+
  |   |
      |
      |
      |
      |
=========`)

	case 5:
		fmt.Println(`
  +---+
  |   |
  O   |
      |
      |
      |
=========`)

	case 4:
		fmt.Println(`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`)

	case 3:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`)

	case 2:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`)

	case 1:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`)

	case 0:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`)

	}
}





