package main

import "fmt"

func cek(angka int) bool {
	result := true
	if angka != 0 {
		for i := 2; i*i < angka; i++ {
			if angka%i == 0 {
				result = false
			}
		}
	} else {
		result = false
	}
	return result

}

func primeX(number int) int {
	tampung := 0
	angka := 1

	for tampung < number {
		angka++
		if cek(angka) {
			tampung++
		}
	}
	return angka
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
