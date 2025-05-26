package dto

type UserDto struct {
	Id 			uint32
	Telegram_id string
	Username	string
}

type CreateUserDto struct {
	Telegram_id string
	Username	string
	Password 	string
}

type UpdateUserDto struct {
	Id 			uint32
	Telegram_id string
	Username	string
	Password 	string
}