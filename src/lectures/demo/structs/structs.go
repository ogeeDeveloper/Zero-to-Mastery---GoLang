package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	// Create passengers
	billy := Passenger{"Billy", 12, true}

	fmt.Println(billy)

	var (
		jim  = Passenger{Name: "Jim", TicketNumber: 3}
		lane = Passenger{Name: "Lane", TicketNumber: 4}
	)

	fmt.Println(jim, lane)

	var tom Passenger
	tom.Name = "Tom"
	tom.TicketNumber = 5
	tom.Boarded = true
	fmt.Println(tom)

	billy.Boarded = false
	fmt.Println(billy)
}
