package main

import (
	"fmt"
)

func minCostFrog(jumps []int) int {
	n := len(jumps)
	if n <= 1 {
		return 0
	}

	minCost := make([]int, n)

	minCost[0] = 0
	minCost[1] = abs(jumps[1] - jumps[0])

	for i := 2; i < n; i++ {
		minCost[i] = min(minCost[i-1]+abs(jumps[i]-jumps[i-1]), minCost[i-2]+abs(jumps[i]-jumps[i-2]))
	}

	return minCost[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minCostFrog([]int{10, 10, 40, 20}))         // Output: 30
	fmt.Println(minCostFrog([]int{30, 10, 60, 10, 60, 50})) // Output: 40
}
