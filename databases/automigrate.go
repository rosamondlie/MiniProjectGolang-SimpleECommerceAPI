package databases

import (
	"final-project/configs"
	"final-project/models"
	"fmt"
)

func AutoMigrate() {
	err := configs.DB.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Product{},
	)
	if err != nil {
		errorLog := fmt.Sprintf("Gagal Auto Migrate: %s", err.Error())
		panic(errorLog)
	}
}
