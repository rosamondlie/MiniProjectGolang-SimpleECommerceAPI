package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string  `json:"nama"`
	Email    string  `json:"email"`
	NoHP     *string `json:"no_hp"`
	Status   bool    `json:"status" gorm:"default:true"`
}
