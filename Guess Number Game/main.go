package main

import (
	"fmt"
	"math/rand"
)

func main() {

	min,max:=1,100
	number := rand.Intn(max-min+1) + min
	try := 10

	fmt.Println("I picked a number between 1 and 100, guess it:")
	for try > 0{

		var guess int
		fmt.Scanln(&guess)

		if guess > number{
			fmt.Println("Go down!")
			try--
			if(try==0){
				fmt.Println("No tries left....")
			}
		}else if guess < number{
			fmt.Println("Go Up!")
			try--
			if(try==0){
				fmt.Println("No tries left....")
			}
		}else{
			fmt.Println("You guessed the number!")
			break
		}
	}
}



