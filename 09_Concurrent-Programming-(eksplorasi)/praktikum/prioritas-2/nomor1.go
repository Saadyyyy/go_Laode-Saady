package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	// Membagi teks menjadi beberapa bagian
	textParts := strings.Split(text, " ")

	// Membuat map untuk menghitung frekuensi huruf
	freqMap := make(map[rune]int)

	// Membuat mutex untuk mengamankan akses ke map
	var mu sync.Mutex

	// Membuat WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Fungsi goroutine untuk menghitung frekuensi huruf dalam satu bagian teks
	countLetter := func(part string) {
		defer wg.Done()

		// Menghitung frekuensi huruf dalam bagian teks
		partFreq := make(map[rune]int)
		for _, char := range part {
			if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
				partFreq[char]++
			}
		}

		// Mengunci (Lock) mutex sebelum mengakses map frekuensi huruf
		mu.Lock()
		defer mu.Unlock() // Memastikan mutex selalu dilepaskan setelah selesai

		// Menggabungkan hasil perhitungan dari bagian ke map frekuensi huruf total
		for char, count := range partFreq {
			freqMap[char] += count
		}
	}

	// Menjalankan goroutine untuk menghitung frekuensi huruf dalam setiap bagian teks
	for _, part := range textParts {
		wg.Add(1)
		go countLetter(part)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	// Mencetak frekuensi huruf
	for char, count := range freqMap {
		fmt.Printf("Huruf '%c' muncul sebanyak %d kali\n", char, count)
	}
}
