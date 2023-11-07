package database

import (
	"github.com/Thund3rD3v/SuperGuardian/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Debug Only
	db.AutoMigrate(&structs.Member{})

	return db
}
