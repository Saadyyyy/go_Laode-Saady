package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		// fmt.Println(i)
		tampung := ""

		if i%3 == 0 {
			tampung += "Fizz"
		}

		if i%5 == 0 {
			tampung += "Buzz"
		}

		if i%3 == 0 || i%5 == 0 {
			fmt.Println(tampung)
		} else {
			fmt.Println(i)
		}
	}
}
