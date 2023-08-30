package main

import "fmt"

func fibonacci(number int) int {
	var result int
	if number == 0 {
		result = 0
	} else if number == 1 {
		result = 1
	} else {
		a, b := 0, 1
		for i := 2; i <= number; i++ {
			a, b = b, a+b
			result = b
		}
	}
	return result
}

func main() {

	// fmt.Println(fibonacci(0)) // 0

	// fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	// fmt.Println(fibonacci(10)) // 55

	// fmt.Println(fibonacci(12)) // 144

}
