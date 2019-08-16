package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggyBank float64
	money := []float64{0.05, 0.10, 0.25}
	for piggyBank <= 20.00 {
		piggyBank += money[rand.Intn(3)]
		fmt.Printf("%5.2f\n", piggyBank)
	}
}
