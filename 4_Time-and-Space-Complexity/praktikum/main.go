package main

import "fmt"

func main() {
	// fmt.Println("Hello world!")
	fmt.Println(number(1000000007))
	fmt.Println(number(25))
	fmt.Println(number(13))
	fmt.Println(number(17))
	fmt.Println(number(20))
	fmt.Println(number(35))
}

func number(num int) bool {

	var result bool
	for i := 2; i < num/2; i++ {

		if num%i == 0 {
			result = false
			break
		} else {
			result = true
		}
	}
	// if result == false{
	// 	fmt.Println("Bukan Bilangan Primer")
	// }else{
	// 	fmt.Println("Bilangan Primer")
	// }
	return result
}
