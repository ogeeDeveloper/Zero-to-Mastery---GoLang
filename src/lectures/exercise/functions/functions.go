//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetPerson(firstName string) {
	fmt.Println("Hello, how are you doing ", firstName)
}

func sumThreeNums(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func printNum(numberGiven int) int {
	return numberGiven
}

// func printTwoNumbers(oneNum, twoNum int) int {
// 	return oneNum, twoNum
// }

func main() {
	sumThreeNums(2, 5, 6)
	providedNum := printNum(3)
	fmt.Println("Number provvided is", providedNum)
	greetPerson("Mary")
}
