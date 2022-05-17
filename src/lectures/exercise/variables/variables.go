//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favouriteColor = "blue"
	fmt.println("My favourite color is ", favouriteColor)

	birtYear, birthAge := 4040, 300
	fmt.println("I was bor in ", birthYear, "and m age is", birthAge)

	var (
		firstInitial = "J"
		lastInitial  = "P"
	)
	fmt.println("First Initial= ", firstInitial, "Last Initial= ", lastInitial)

	var ageInDays int
	ageInDays = birthAge * 365
	fmt.println("My bith age in days is ", ageInDays)

}
