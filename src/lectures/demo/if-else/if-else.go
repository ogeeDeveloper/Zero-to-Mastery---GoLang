package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 7, 6, 9

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 is greater than quiz 2")
	} else if quiz2 > quiz1 {
		fmt.Println("Quiz 2 is greater than quiz 1")
	} else {
		fmt.Println("Both quizes have the same course")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Your scores are acceptable")
	}
}
