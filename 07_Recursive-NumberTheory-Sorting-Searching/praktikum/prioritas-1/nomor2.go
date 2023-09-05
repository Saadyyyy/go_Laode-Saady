package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	cek := make(map[string]int)
	var result []pair
	for _, v := range items {
		if _, isExits := cek[v]; isExits {
			cek[v]++
		} else {
			cek[v] = 1
		}
	}

	for name, count := range cek {
		result = append(result, pair{name, count})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].count < result[j].count
	})

	return result
}

func main() {

	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}
}
