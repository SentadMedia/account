package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
)

var _ repository.Account = (*AccountSQL)(nil)

// AccountSQL accesses User information in user table through SQL.
type AccountSQL struct {
	db *gorm.DB
}

// NewAccountSQL creates UserSQL
func NewAccountSQL(db *gorm.DB) AccountSQL {
	return AccountSQL{db: db}
}

// RegisterAccount Actually writes a user into DB
func (p AccountSQL) RegisterAccount(username, email, password string) (entity.Account, error) {
	account := entity.Account{
		Username: username,
		Email:    email,
		Password: password,
	}

	if dbc := p.db.Create(&account); dbc.Error != nil {
		return entity.Account{}, dbc.Error
	}

	return account, nil
}
