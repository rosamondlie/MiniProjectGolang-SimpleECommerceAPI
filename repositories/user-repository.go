package repositories

import (
	"final-project/configs"
	"final-project/models"
)

func CreateUser(user *models.User) error {
	err := configs.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func ListUser()(users []models.User, err error) {
	err = configs.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(user *models.User) error {
	err := configs.DB.Save(user).Error
	if err != nil{
		return err
	}
	return nil
}

func GetUserByID(id uint)(user *models.User, err error){
	err = configs.DB.First(&user,id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id uint)error{
	err := configs.DB.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil	
}