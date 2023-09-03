package main

import "fmt"

func Frog(jumps []int) int {
	var result int
	for i := 0; i < len(jumps); i++ {
		for j := 0; j < i; j++ {
			result += j
		}
	}
	return result
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
