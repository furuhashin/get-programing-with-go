package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	//a := 'c'
	//b := a + 3
	//fmt.Printf("%c", b)
	for _, t := range cipherText {
		// A 65 C 67 G71
		// A 0  C 2  G 6
		// C-G = -4 = 26-4
		c := uint8(t - 'A')          // 0 1 2 ...25
		k := keyword[keyIndex] - 'A' // 0 1 2 ...25

		c = (c-k+26)%26 + 'A' // (2-6+26)%26 = 22  22 + 'A' = 'W'
		message += string(c)
		// increment keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)
}
