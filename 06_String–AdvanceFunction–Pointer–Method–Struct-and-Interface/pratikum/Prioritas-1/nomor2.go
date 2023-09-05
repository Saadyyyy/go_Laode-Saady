package main

import "fmt"

type student struct {
	name  string
	score int
}

func numScore(s []student) float32 {
	tampungScore := 0
	for _, v := range s {
		tampungScore += v.score
	}
	rataRata := float32(tampungScore) / float32(len(s))

	return rataRata
}

func main() {
	s := []student{
		student{name: "Laode", score: 90},
		student{name: "Saady", score: 80},
		student{name: "Rian", score: 40},
		student{name: "Agung", score: 30},
		student{name: "Ira", score: 70},
	}
	rataRata := numScore(s)
	scoreTertinggi := s[0].score
	scoreTerrendah := s[0].score
	for _, v := range s {
		if v.score > scoreTertinggi {
			scoreTertinggi = v.score
		}
		if v.score < scoreTerrendah {
			scoreTerrendah = v.score
		}
	}
	var namaTertinggi string
	var namaTerendah string

	for _, v := range s {
		if v.score == scoreTertinggi {
			namaTertinggi = v.name
		}
		if v.score == scoreTerrendah {
			namaTerendah = v.name
		}
	}
	fmt.Printf("Nilai rata -rata : %.2f\nNilai Tertinggi : %s : %d\nNilai Terendah : %s : %d", rataRata, namaTertinggi, scoreTertinggi, namaTerendah, scoreTerrendah)
}
