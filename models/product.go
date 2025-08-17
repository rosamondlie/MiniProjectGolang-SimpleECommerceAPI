package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Nama            string  `json:"nama"`
	Harga           int     `json:"harga"`
	Stok            int     `json:"stok"`
	PenanggungJawab string  `json:"penanggung_jawab"`
	Photo           *string `json:"photo"`
}
