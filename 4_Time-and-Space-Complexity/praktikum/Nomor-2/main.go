package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	fmt.Println(pangkat(2,3))
	 fmt.Println(pangkat(5, 3))  // 125
   fmt.Println(pangkat(10, 2)) // 100
   fmt.Println(pangkat(2, 5))  // 32
   fmt.Println(pangkat(7, 3))  // 343

}

func pangkat(x,n int)int{
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}