package entity

import (
	"github.com/jinzhu/gorm"
)

// Role Data type
type Role struct {
	gorm.Model
	AccountID   uint
	Name        string       `gorm:"type:varchar(255);default:'roleName'"`
	Permissions []Permission `gorm:"many2many:role_permissions;PRELOAD;"`
}
