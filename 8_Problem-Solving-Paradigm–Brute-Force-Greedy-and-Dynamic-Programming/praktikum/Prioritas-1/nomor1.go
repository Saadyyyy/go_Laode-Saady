package main

import (
	"fmt"
)

func bilanganBiner(n int) []string {
	result := []string{}
	for i := 0; i <= n; i++ {
		binaryString := fmt.Sprintf("%b", i)
		result = append(result, binaryString)
	}
	return result
}

func main() {
	fmt.Println(bilanganBiner(3))
}

/*
Diberi bilangan bulat n, kembalikan array ans dengan panjang n + 1 sehingga untuk setiap
i (0 <= i <= n), ans[i] adalah bilangan 1 dalam representasi biner dari i

Input: n = 2
Output: [0,1,10]

Input: n = 3
Output: [0,1,10, 11]

input : n =6
output : [0,1,10,11,100,110,111*/
