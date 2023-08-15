package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++{
		if i % 2 == 0 {
			fmt.Println("Angka genap :")
		}
		if i % 2 == 1{
			fmt.Println("Angka ganjil :")
		}

		fmt.Println(i)
	}
}