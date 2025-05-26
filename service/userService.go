package service

import (
	"log"
	"notification_telegram_bot/dto"
	"notification_telegram_bot/mapper"
	"notification_telegram_bot/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) GetAll() []dto.UserDto {
	users, err := service.repository.GetAll()
	var dtos []dto.UserDto

	if err != nil {
		log.Println(err)
		return nil
	}

	for _, user := range users {
		dto := mapper.ToDto(&user)
		dtos = append(dtos, *dto)
	}

	return dtos
}

func (service *UserService) GetById(id uint32) *dto.UserDto {
	user, err := service.repository.GetById(id)

	if err != nil {
		log.Println(err)
		return nil
	}

	dto := mapper.ToDto(user)

	return dto
}

func (service *UserService) Create(dto *dto.CreateUserDto) *dto.UserDto {
	user := mapper.ToModel(dto)
	createdUser, err := service.repository.Create(user)

	if err != nil {
		log.Println(err)
		return nil
	}

	createdUserDto := mapper.ToDto(createdUser)

	return createdUserDto
}

func (service *UserService) Update(dto dto.UpdateUserDto) *dto.UserDto {
	id := dto.Id
	user, err := service.repository.GetById(id)

	if err != nil {
		log.Println(err)
		return nil
	}

	updatedUser := mapper.Update(&dto, user)
	newUpdatedUser, err := service.repository.Update(updatedUser)

	if err != nil {
		log.Println(err)
		return nil
	}

	updatedUserDto := mapper.ToDto(newUpdatedUser)

	return updatedUserDto
}

func (service *UserService) Delete(id uint32) {
	service.repository.Delete(id)
}