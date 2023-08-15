package main

import "fmt"

func main() {
	palindrom := "aan"
	for i := 0; i < len(palindrom); i++ {
		a := palindrom[i]
		rumus := palindrom[len(palindrom)-i-1]

		if a != rumus {
			fmt.Println("Bukan Palindrom")
			break
		} else {
			fmt.Println("Ini kata Palindrom")
			break
		}
	}
}
