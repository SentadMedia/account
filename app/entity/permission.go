package entity

import (
	"github.com/jinzhu/gorm"
)

// Permission Data type
type Permission struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);default:'permissionName'"`
	ContentType string `gorm:"type:varchar(255);not null;"`
	CodeName    string `gorm:"type:varchar(255);not null;"`
}
