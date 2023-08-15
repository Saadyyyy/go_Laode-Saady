package main

import "fmt"

func main() {
	bintang := 10
	for i := 1; i <= bintang; i++ {
		// fmt.Println(i)
		tampung := ""
		for b := i; b < bintang; b++ {
			tampung += " "
		}

		for a := 1; a <= i; a++ {
			tampung += "* "
		}
		fmt.Println(tampung)

	}
}
