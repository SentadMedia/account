package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/sentadmedia/account/app/usecase/repository"
)

var _ repository.Account = (*AccountSQL)(nil)

// AccountSQL accesses User information in user table through SQL.
type AccountSQL struct {
	db *gorm.DB
}

func (a AccountSQL) GetAcccount() string {
	fmt.Print(a.db)
	return "kawtar elouaraini"
}

// NewUserSQL creates UserSQL
func NewAccountSQL(db *gorm.DB) AccountSQL {
	return AccountSQL{db: db}
}
