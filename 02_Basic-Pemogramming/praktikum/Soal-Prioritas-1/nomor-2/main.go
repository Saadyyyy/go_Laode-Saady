package main

import "fmt"

func main() {
	number := 100
	for i := 0; i <= number; i++{
		if i % 2 == 0 {
			fmt.Println(i,"Angka genap ")
		}
		if i % 2 == 1{
			fmt.Println(i,"Angka ganjil ")
		}

		// fmt.Println(i)
	}
}