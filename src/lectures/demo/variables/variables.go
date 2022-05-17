package main

import "fmt"

func main() {
	var myName = "John"
	fmt.Println("My name is", myName)

	var name string = "Jane"
	fmt.Println("name=", name)

	// Create and assign
	userName := "admin"
	fmt.Println("Username = ", userName)

	// Unassigned variable
	var sum int
	fmt.Println("The sum is: ", sum)

	// Compound assignment
	part1, other := 1, 5
	fmt.Println("Part1 is", part1, "Other is: ", other)
}
