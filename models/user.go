package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string  `json:"nama"`
	Email    string  `json:"email"`
	NoHP     *string `json:"no_hp"`
	Password string  `json:"-"`
	Status   bool    `json:"lulus" gorm:"default:0"`
}
