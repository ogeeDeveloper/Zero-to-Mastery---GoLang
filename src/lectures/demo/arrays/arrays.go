package main

import "fmt"

// Create a structure to holds information about rooms
type Room struct {
	name    string
	cleaned bool
}

// Create a function to keep track opf all rooms to see if they have been cleaned
// Array will hold 4 rooms
func checkCleaned(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i] // Create a copy of the rooms array

		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, " is dirty and should be cleaned")
		}
	}
}

func main() {
	// Create an array of rooms
	rooms := [...]Room{
		{name: "Cafeteria"},
		{name: "Office"},
		{name: "Reception"},
		{name: "HR"},
	}

	checkCleaned(rooms)
	fmt.Println("Performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleaned(rooms)
}
