package main

import (
	"fmt"
)

func SolveEquations(a, b, c int) (x, y, z int, solution bool) {
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000; y++ {
			for z := 1; z <= 10000; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+2*z == c {
					return x, y, z, true
				}
			}
		}
	}
	return 0, 0, 0, false
}

func main() {
	a, b, c := 6, 6, 14
	x, y, z, solution := SolveEquations(a, b, c)
	if solution {
		fmt.Printf("Solusi ditemukan: x=%d, y=%d, z=%d\n", x, y, z)
	} else {
		fmt.Println("Tidak ada solusi")
	}
}
