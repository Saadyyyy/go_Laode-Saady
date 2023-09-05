package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// your code here
	var result string
	tampung := len(a)
	if len(a) > len(b) {
		tampung = len(b)
	}

	for i := 0; i < tampung; i++ {
		if a[i] == b[i] {
			result += string(a[i])
		} else {
			result += string(b[i])
		}
	}

	return result
}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) //AKA

	fmt.Println(Compare("KANGOORO", "KANG")) //KANG

	fmt.Println(Compare("KI", "KIJANG")) //KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

	fmt.Println(Compare("ILALANG", "ILA")) //ILA

}
