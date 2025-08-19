package seeders

import (
	"errors"
	"final-project/configs"
	"final-project/models"
	"log"

	"gorm.io/gorm"
)

func SeederProducts() {
	err := configs.DB.First(&models.Product{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		products := []models.Product{
			{
				Nama:            "Laptop ASUS Vivobook",
				Harga:           8500000,
				Stok:            10,
				PenanggungJawab: "Budi Santoso",
				Photo:           nil,
			},
			{
				Nama:            "Headset Logitech G733",
				Harga:           2500000,
				Stok:            15,
				PenanggungJawab: "Budi Santoso",
				Photo:           nil,
			},
			{
				Nama:            "Smart TV Samsung 43 Inch",
				Harga:           5500000,
				Stok:            7,
				PenanggungJawab: "Budi Santoso",
				Photo:           nil,
			},
		}

		for _, product := range products {
			err := configs.DB.Create(&product).Error
			if err != nil {
				log.Printf("Gagal insert product %s: %v", product.Nama, err)
			}
		}
	}
}
