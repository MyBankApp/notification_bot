package db

import (
	"log"
	"notification_telegram_bot/model"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func Connect(pathToDb string) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(pathToDb), &gorm.Config{})

		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		sqlDB, err := db.DB()

		if err != nil {
            log.Fatalf("Failed to get DB: %v", err)
        }

        sqlDB.SetMaxOpenConns(1)
		err = db.AutoMigrate(&model.User{})

		if err != nil {
			panic("Migration failed: " + err.Error())
		}

		dbInstance = db
	})

	return dbInstance
}
