package model

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Judul    string `json:"judul"  form:"judul"`
	Penulis  string `json:"penulis"  form:"penulis"`
	Penerbit string `json:"penerbit"  form:"penerbit"`
}
