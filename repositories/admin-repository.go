package repositories

import (
	"final-project/configs"
	"final-project/models"
)

func FindByEmail(email string) (*models.Admin, error) {
	var user models.Admin
	if err := configs.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByPhone(phone string) (*models.Admin, error) {
	var user models.Admin
	if err := configs.DB.Where("no_hp = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
