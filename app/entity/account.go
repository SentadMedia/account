package entity

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	ID       int64  `gorm:"primary_key"`
	Username string `gorm:"type:varchar(255);unique_index;not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
