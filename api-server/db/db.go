package db

import (
	"api-server/config"
	"api-server/types"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

func GetDb(c config.Config, l *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=Europe/Berlin",
		c.DbHost(),
		c.DbUser(),
		c.DbPassword(),
		c.DbName(),
		c.DbPort(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		error := fmt.Sprintf(
			"Db: failed to connect to database '%s'",
			strings.ReplaceAll(dsn, c.DbPassword(), "hidden"),
		)
		l.Error(error)
		panic(error)
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
