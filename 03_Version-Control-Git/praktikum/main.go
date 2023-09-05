package main

import (
	"fmt"
)

func main() {
	var palindrom string
	fmt.Printf("Input Disini : ")
	fmt.Scan(&palindrom)
	b := true
	for i := 0; i < len(palindrom)/2; i++ {
		a := palindrom[i]
		rumus := palindrom[len(palindrom)-i-1]

		if a != rumus {
			b = false
			break
		}
	}
	if b {
		fmt.Println("Ini kata Palindrom")
		// break
	} else {
		fmt.Println("Bukan Palindrom")
		// break
	}
}
