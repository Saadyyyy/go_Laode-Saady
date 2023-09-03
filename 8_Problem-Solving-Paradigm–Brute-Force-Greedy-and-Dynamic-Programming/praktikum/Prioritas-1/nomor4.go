package main

func SimpleEquations(a, b, c int) {
	x := 1
	y := 2
	z := 3
	a = x + y + z
	b = x * y * z
	c = x ^ 2 + y ^ 2 + z*2

}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
