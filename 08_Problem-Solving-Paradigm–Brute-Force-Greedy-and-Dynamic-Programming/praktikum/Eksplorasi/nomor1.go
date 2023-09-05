package main

import "fmt"

func angkaRomawi(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanNumeral := ""
	i := 0

	for num > 0 {
		for num >= val[i] {
			romanNumeral += syb[i]
			num -= val[i]
		}
		i++
	}

	return romanNumeral
}

func main() {
	fmt.Println(angkaRomawi(4))    // output IV
	fmt.Println(angkaRomawi(9))    // output IX
	fmt.Println(angkaRomawi(23))   // output XXIII
	fmt.Println(angkaRomawi(2021)) // output MMXXI
	fmt.Println(angkaRomawi(1646)) // output MDCXLVI
}
