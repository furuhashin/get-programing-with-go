package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	lightYears := distance / lightSpeed / secondsPerDay / daysPerYear
	fmt.Println(lightYears)
}
