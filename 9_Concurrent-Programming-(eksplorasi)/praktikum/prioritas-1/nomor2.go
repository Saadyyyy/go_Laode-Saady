package main

import (
	"fmt"
	"time"
)

func keliapatanAngka2(x int, ch chan int) {
	for i := 0; ; i++ {
		if i%x == 0 {
			ch <- i
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	Kelipatan := 3
	ch := make(chan int)

	go keliapatanAngka2(Kelipatan, ch)

	for {
		select {
		case kelipatan := <-ch:
			fmt.Printf("%d Adalah kelipatan dari %d \n", kelipatan, Kelipatan)
			if kelipatan >= 100 {
				fmt.Printf("Kelipatan %d Selesai\n", Kelipatan)
				return
			}
		}
	}
}
