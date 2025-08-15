package repositories

import (
	"final-project/configs"
	"final-project/models"
)

func ListProduct() (product []models.Product, err error){
	err = configs.DB.Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func GetProductByID(id uint)(product *models.Product, err error){
	err = configs.DB.First(&product,id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(product models.Product) error {
	err := configs.DB.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id uint, product *models.Product) error {
	err := configs.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id uint) error {
	err := configs.DB.Delete(&models.Product{}, id).Error
	if err != nil {
		return err
	}

	return nil
}