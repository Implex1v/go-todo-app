package types

import (
	"fmt"
	"go.uber.org/fx"
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
		User{},
	)
	if err != nil {
		panic("Failed to migrate types")
	}

	db.Create(&User{Password: "123456", Username: "Longcat", Email: "longcat@cat.long"})

	return db
}

var Module = fx.Options(
	fx.Provide(GetDb),
)
