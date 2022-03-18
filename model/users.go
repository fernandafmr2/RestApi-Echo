package model

import "RestApi-Echo/config"

type Users struct {
	Email       string `json:"email" form:"email"`
	Name        string `json:"name" form:"name" gorm:"primaryKey"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

func (user *Users) CreateUser() error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateUser(name string) error {
	if err := config.DB.Model(&Users{}).Where("name = ?", name).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByName(name string) (Users, error) {
	var user Users
	result := config.DB.Where("name = ?", name).First(&user)
	return user, result.Error
}

func GetAll(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Where("name LIKE ? OR name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)

	return users, result.Error
}
