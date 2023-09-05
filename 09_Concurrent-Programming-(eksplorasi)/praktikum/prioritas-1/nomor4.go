package main

import (
	"fmt"
	"sync"
)

func main() {
	// Membuat mutex untuk mengamankan akses ke daftar (slice)
	var mu sync.Mutex

	// Inisialisasi daftar
	var myList []int

	// Membuat WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Jumlah goroutine yang akan dijalankan
	totalGoroutines := 5

	// Fungsi goroutine untuk menambahkan elemen ke daftar
	goroutineFunc := func(i int) {
		defer wg.Done()

		// Mengunci (Lock) mutex sebelum mengakses daftar
		mu.Lock()
		defer mu.Unlock() // Memastikan mutex selalu dilepaskan setelah selesai

		// Menambahkan elemen ke daftar
		myList = append(myList, i)
		fmt.Printf("Goroutine %d: Menambahkan %d ke daftar\n", i, i)
	}

	// Menjalankan beberapa goroutine
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go goroutineFunc(i)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	// Mencetak isi akhir daftar
	fmt.Printf("Isi Daftar Akhir: %v\n", myList)
}
