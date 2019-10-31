package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 233.0
	fmt.Printf("233º K is %.2fº C\n", kelvinToCelsius(kelvin))
	fmt.Printf("0º K is %.2fº F\n", kelvinToFahrenheit(0))
}
