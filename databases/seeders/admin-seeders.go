package seeders

import (
	"errors"
	"final-project/configs"
	"final-project/models"
	"final-project/utils"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SeederAdmin() {
	err := configs.DB.First(&models.Admin{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var admin []models.Admin

		hashPassword, err := utils.HashPassword("admin123")
		if err != nil {
			fmt.Println("error :", err.Error())
			return
		}

		admin = append(admin,
			models.Admin{
				Nama:     "Budi Santoso",
				Email:    "admin@example.com",
				NoHP:     utils.StringPtr("08121231234"),
				Password: hashPassword,
			})

		if err := configs.DB.Create(&admin).Error; err != nil {
			log.Println("Gagal membuat admin seeder:", err)
		} else {
			log.Println("Seeder admin berhasil dibuat!")
		}
	}
}
