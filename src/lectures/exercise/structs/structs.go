//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

//* Create a rectangle structure containing its coordinates
type rectangle struct {
	length, width int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
func calculateArea(length, width int) int {
	return length * width
}

func main() {
	rectangle.length = 4
	rectangle.width = 6
	result := calculateArea(rectangle.length, rectangle.width)
	fmt.Println(result)
}
