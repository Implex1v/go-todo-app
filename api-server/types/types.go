package types

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	title   string
	dueDate time.Time
}
