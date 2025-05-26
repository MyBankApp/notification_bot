package repository

import (
	"notification_telegram_bot/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *UserRepository {
	return &UserRepository{Db: Db}
}

func (repository *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	result := repository.Db.Find(&users)
	return users, result.Error
}

func (repository *UserRepository) GetById(id uint32) (*model.User, error) {
	var user *model.User
	result := repository.Db.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (repository *UserRepository) Create(user *model.User) (*model.User, error) {
	result := repository.Db.Create(user)
	return user, result.Error
}

func (repository *UserRepository) Update(user *model.User) (*model.User, error) {
	result := repository.Db.Save(user)
	return user, result.Error
}

func (repository *UserRepository) Delete(id uint32) {
	repository.Db.Where("id = ?", id).Delete(&model.User{})
}