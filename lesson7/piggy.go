package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggyBank int
	money := []int{5, 10, 25}
	for piggyBank <= 2000 {
		piggyBank += money[rand.Intn(3)]

		dollars := piggyBank / 100
		cents := piggyBank % 100
		fmt.Printf("$%d.%02d\n", dollars, cents)
	}
}
