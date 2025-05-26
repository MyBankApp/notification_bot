package model

import "gorm.io/gorm"

type User struct {
	Id 			uint32 	`gorm:"primaryKey"`
	Telegram_id	string	`gorm:"unique"`
	Username	string	`gorm:"unique"`
	Password	string	`gorm:"unique"`
	gorm.Model
}
