package main

import "fmt"

func double(x int) int {
	return x + x
}

func sum(num1, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println("double =", double(6))
	fmt.Println("Sum is ", sum(1, 4))
}
