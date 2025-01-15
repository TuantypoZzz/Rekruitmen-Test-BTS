package entity

import (
	"github.com/google/uuid"
)

type Checklist struct {
	Id       uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36)"`
	Title    string    `gorm:"column:title"`
	Username string    `gorm:"column:username"`
}
