package mapper

import (
	"notification_telegram_bot/dto"
	"notification_telegram_bot/model"

	"github.com/jinzhu/copier"
)

func ToModel(dto *dto.CreateUserDto) *model.User {
	user := model.User{}
	copier.Copy(&user, &dto)
	return &user
}

func Update(dto *dto.UpdateUserDto, user *model.User) *model.User {
	copier.Copy(&user, &dto)
	return user
}

func ToDto(user *model.User) *dto.UserDto {
	dto := dto.UserDto{}
	copier.Copy(&dto, &user)
	return &dto
}
