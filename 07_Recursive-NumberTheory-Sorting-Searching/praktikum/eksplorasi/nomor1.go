package main

import "fmt"

func MaxSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sum := arr[0]
	result := arr[0]

	for _, v := range arr[1:] {
		// currentSum = max(value, (currentSum + value))
		// resultSum = max(resultSum, currentSum)

		tampung := sum + v
		if v > tampung {
			sum = v
		} else if v < tampung {
			sum = tampung
		}

		if sum > result {
			result = sum
		}
	}
	return result
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12
}
