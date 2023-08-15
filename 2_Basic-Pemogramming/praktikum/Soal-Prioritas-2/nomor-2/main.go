package main

import "fmt"

func main() {

	nilai := 30

	for i := 1; i <= nilai; i++ {
		// fmt.Println(i)
		if nilai%i == 0 {
			fmt.Println(i)
		}
	}
	// fmt.Println(6*5)
}
