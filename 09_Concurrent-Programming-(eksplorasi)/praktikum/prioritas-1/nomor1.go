package main

import (
	"fmt"
	"time"
)

func keliapatanAngka(x int) {
	for i := 0; ; i++ {
		if i%x == 0 {
			fmt.Printf("%d Adalah kelipatan dari %d \n", i, x)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	Kelipatan := 3
	go keliapatanAngka(Kelipatan)

	time.Sleep(100 * time.Second)
	fmt.Printf("Kelipatan %d Selesai", Kelipatan)
}
