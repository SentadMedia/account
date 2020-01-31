package entity

import (
	"github.com/jinzhu/gorm"
)

// Account Data type
type Account struct {
	gorm.Model
	Username    string       `gorm:"UNIQUE_INDEX;"`
	FirstName   string       `gorm:"type:varchar(255);default:'firstName'"`
	LastName    string       `gorm:"type:varchar(255);default:'lastName'"`
	Email       string       `gorm:"type:varchar(255);not null;unique;"`
	Password    string       `gorm:"not null;" json:"-"`
	IsSuperuser bool         `gorm:"default:false;"`
	Permissions []Permission `gorm:"many2many:account_permissions;PRELOAD;"`
	Role        Role         `gorm:"foreignkey:AccountID;PRELOAD;"`
}
