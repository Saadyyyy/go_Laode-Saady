// Buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”. Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz …….

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		// fmt.Println(i)
		tampung := ""

		if i%3 == 0 {
			tampung += "Fizz"
		}

		if i%5 == 0 {
			tampung += "Buzz"
		}

		if i%3 == 0 || i%5 == 0 {
			fmt.Println(tampung)
		} else {
			fmt.Println(i)
		}
	}
}
