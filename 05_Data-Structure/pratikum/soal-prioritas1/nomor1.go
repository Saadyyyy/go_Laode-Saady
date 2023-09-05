package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	var result []string
	mapping := make(map[string]bool)

	for _, va := range arrayA {
		result = append(result, va)
		mapping[va] = true
	}
	for _, vb := range arrayB {
		if mapping[vb] == false {
			result = append(result, vb)
		}
	}

	return result
}

func main() {

	// Test cases

	// fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	// fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	// fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	// fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	// fmt.Println(ArrayMerge([]string{}, []string{}))

	// []

}
