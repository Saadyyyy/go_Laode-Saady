package main

import (
	"fmt"
)

func segitigaPascal(numRows int) [][]int {
	result := [][]int{}
	for i := 0; i < numRows; i++ {
		baris := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				baris = append(baris, 1)
			} else {
				kiri := result[i-1][j-1]
				kanan := result[i-1][j]
				baris = append(baris, kiri+kanan)
			}
		}
		result = append(result, baris)
	}
	return result
}

func main() {
	numRows := 5
	segitigaPascal := segitigaPascal(numRows)

	fmt.Println(segitigaPascal)
}
