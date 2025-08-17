package repositories

import (

	"final-project/configs"
	"final-project/models"
)

func ListProduct() (product []models.Product, err error) {
	err = configs.DB.Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func GetProductByID(id uint) (product *models.Product, err error) {
	err = configs.DB.First(&product, id).Error
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
	err := configs.DB.Save(product).Error
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

func CountProducts() (int64, error) {
	var count int64
	err := configs.DB.Model(&models.Product{}).Count(&count).Error
	return count, err
}

func CountAvailableProducts() (int64, error) {
	var count int64
	err := configs.DB.Model(&models.Product{}).
		Where("stok > ?", 0).
		Count(&count).Error
	return count, err
}

func GetLatestProducts(limit int) ([]models.Product, error) {
	var products []models.Product
	err := configs.DB.Order("created_at desc").Limit(limit).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetAvailableProducts() ([]models.Product, error) {
	var products []models.Product
	err := configs.DB.Model(&models.Product{}).
		Where("stok > ?", 0).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

