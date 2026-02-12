package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// Iから始まるものはインターフェース
// インターフェースはメソッドの一覧
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB	
}

func NewUserRepositry(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByEmail (user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}
func (ur *UserRepository) CreateUser (user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}