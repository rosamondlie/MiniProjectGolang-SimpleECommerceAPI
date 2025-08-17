package seeders

import (
	"errors"
	"final-project/configs"
	"final-project/models"

	"gorm.io/gorm"
)

func SeederAdmin() {
	err := configs.DB.First(&models.Admin{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var admin []models.Admin

		admin = append(admin,
			models.Admin{
				Nama: "Budi Santoso",
			},
			models.Admin{
				Nama: "Monmond",
			})

		configs.DB.Create(&admin)
	}
}

