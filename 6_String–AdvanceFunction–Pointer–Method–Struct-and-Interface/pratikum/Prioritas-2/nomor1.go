package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	for _, v := range input {
		intv := (int(v)+offset-97)%26 + 97
		result += string(intv)
	}
	return result
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // def

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	// fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
