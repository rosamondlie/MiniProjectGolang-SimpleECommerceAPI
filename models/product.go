package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Nama		string  `json:"nama"`
	Harga int    `json:"harga"`
	Stok int   `json:"stok"`
	UserID uint `json:"user_id"`
	Photo *string `json:"photo"`
}
