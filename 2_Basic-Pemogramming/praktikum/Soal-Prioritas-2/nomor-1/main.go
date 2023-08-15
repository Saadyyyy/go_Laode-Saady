package main

import "fmt"

func main() {
	bintang := 5
	for i := 1; i <= bintang; i++ {
		// fmt.Println(i)
		tampung := ""

		for a := 1; a <= i; a++ {
			tampung += "* "
			for t := i; t < a; t++ {
				tampung += "* "
			}
		}

		// for t := i; t > 0; t--{
		// 	tampung += "* "
		// }
		fmt.Println(tampung)
	}
}
