package main

import (
	"fmt"
)

func main() {
	ca := "L fdph, L vdz, L frqtxhuhg"
	for i, c := range ca {
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c < 'a' {
				c = c + 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c < 'A' {
				c = c + 26
			}
		}
		fmt.Printf("%v %c\n", i, c)
	}
}
