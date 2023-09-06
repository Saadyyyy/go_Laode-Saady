package main

import "fmt"

type kendaraan struct {
	// disini saya merubah dari totalroda = totalRoda
	// totalRoda       int // saya juga mengomentari juga totalRoda karna totalRoda tidak di gunakan
	kecepatanPerJam int // disini saya merubah dari kecepatanperjam = kecepatanPerJam
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

// function ini mungkin susah di baca karna tidak menerapkan Camel case jadi saya ubah dari kecepatankecapatan = kecapatanKecepatan
func (m *mobil) tambahKecepatan(kecepatanBaru int) { // dan saya mengubah parameter dari kecepatanbaru = kecepatanBaru
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	//agar variable baru dari struct mobil lebih simple jadi saya ubah dari mobilcepat = cepat
	cepat := mobil{}
	cepat.berjalan()
	cepat.berjalan()
	cepat.berjalan()
	// saya sedikit menambah fmt.Println agar kita bisa di tampilkan cepat dari kecepatanPerjam
	fmt.Println("Cepat Mobil :", cepat.kecepatanPerJam)

	//agar variable baru dari struct mobil lebih simple jadi saya ubah dari mobillamban = lamban
	lamban := mobil{}
	lamban.berjalan()
	// saya sedikit menambah fmt.Println agar kita bisa di tampilkan lamban dari kecepatanPerjam
	fmt.Println("Lamban Mobil :", lamban.kecepatanPerJam)
}
