package main

import (
	"fmt"
	"math/rand"
)

func main() {
	distance := 62100000
	secondsPerDay := 86400
	company := []string{"Space Adventure", "SpaceX", "Virgin Galactic"}

	fmt.Println("SpaceLine       Days  Trip-type  Price")
	fmt.Println("======================================")
	for i := 0; i < 10; i++ {
		speed := rand.Intn(15) + 16
		duration := distance / speed / secondsPerDay
		price := 20.0 + speed
		trip := ""
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}
		fmt.Printf("%-16v %4v %-10v %4v\n", company[rand.Intn(3)], duration, trip, price)
	}

}
