//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to roll the dice
func rollDice(sides int) int {
	// add one to it becauise intn fuction satrts at 0
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dice, sides := 2, 6

	rolls := 3

	// This outter for loop will be based on the rolls
	for i := 0; i <= rolls; i++ {
		sum := 0

		// Keeps track of the number of dices
		for diceCounter := 0; diceCounter <= dice; diceCounter++ {
			roll := rollDice(sides)
			sum += roll
			fmt.Println("Roll # is: ", i, "Dice # is: ", diceCounter, ":", roll)
		}
		fmt.Println("Total roll is: ", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake eyes")
		case sum == 7:
			fmt.Println("Lucky 7")
		case sum%2 == 0:
			fmt.Println("Even")
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}

}
