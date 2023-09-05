package main

import "fmt"

func main() {
	nilai := 20

	if nilai <= 100 && nilai >= 0 {
		if nilai >= 80 {
			fmt.Println("A")
		} else if nilai >= 65 {
			fmt.Println("B")
		} else if nilai >= 50 {
			fmt.Println("C")
		} else if nilai >= 35 {
			fmt.Println("D")
		} else if nilai >= 0 {
			fmt.Println("E")
		}
	} else {
		fmt.Println("Nilai Invalid")
	}

}
