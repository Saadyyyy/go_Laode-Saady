package main

import "fmt"

func ArrayMarge(ArrayA, ArrayB []string) []string {
	var result []string
	cek := make(map[string]bool)

	for _, v := range ArrayA {
		result = append(result, v)
		cek[v] = true
	}

	for _, vb := range ArrayB {
		if cek[vb] == false {
			result = append(result, vb)
		}
	}
	return result
}

func main() {
	// fmt.Println(ArrayMarge([]string{"King","Devil Jin", "Akuma"},[]string{"eddie","steve", "geese"}))
	fmt.Println(ArrayMarge([]string{"segei", "jin"}, []string{"jin", "bryan"}))
	// fmt.Println(ArrayMarge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// fmt.Println(ArrayMarge([]string{}, []string{"devil jin", "sergei"}))
	// fmt.Println(ArrayMarge([]string{"hwoarang"}, []string{}))
	// fmt.Println(ArrayMarge([]string{}, []string{}))
}
