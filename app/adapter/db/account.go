package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sentadmedia/account/app/adapter/rpc/proto"

	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
)

var _ repository.Account = (*AccountSQL)(nil)

// AccountSQL accesses User information in user table through SQL.
type AccountSQL struct {
	db *gorm.DB
}

func (a AccountSQL) GetAcccount() string {
	return "kawtar.elouaraini@gmail.com"
}

func (p AccountSQL) CreateAccount(req *proto.PostAccountRequest) (entity.Account, error) {
	account := entity.Account{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	p.db.Create(&account)
	return account, nil
}

// NewUserSQL creates UserSQL
func NewAccountSQL(db *gorm.DB) AccountSQL {
	return AccountSQL{db: db}
}
