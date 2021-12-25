package db

import (
	"api-server/types"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	file := "test.db"
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %s", file))
	}

	err = db.AutoMigrate(
		types.User{},
	)
	if err != nil {
		panic("Failed to migrate types")
	}

	db.Create(&types.User{Password: "123456", Username: "Longcat", Email: "longcat@cat.long"})

	return db
}
