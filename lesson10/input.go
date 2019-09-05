package main

import "fmt"

func main() {

	char := "false11111"
	var result bool

	switch char {
	case "true", "yes", "1":
		result = true
	case "false", "no", "0":
		result = false
	default:
		fmt.Println("error")
	}
	fmt.Println(result)

}
