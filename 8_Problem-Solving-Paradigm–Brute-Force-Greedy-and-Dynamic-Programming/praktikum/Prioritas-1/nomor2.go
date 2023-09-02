package main

import "fmt"

func segitigaPascal(input int) [][]int {
	var result [][]int
	for i := 0; i < input; i++ {
		fmt.Print(i)
	}

	return result
}

func main() {
	fmt.Println(segitigaPascal(5))
}
