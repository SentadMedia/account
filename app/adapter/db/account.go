package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
	"github.com/sentadmedia/elf/fw"
	"golang.org/x/crypto/bcrypt"
)

var _ repository.Account = (*AccountSQL)(nil)

// AccountSQL accesses User information in user table through SQL.
type AccountSQL struct {
	db     *gorm.DB
	logger fw.Logger
}

// NewAccountSQL creates UserSQL
func NewAccountSQL(db *gorm.DB, logger fw.Logger) AccountSQL {
	return AccountSQL{db: db, logger: logger}
}

// RegisterAccount Actually writes a user into DB
func (p AccountSQL) RegisterAccount(account *entity.Account, roleID int32) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		p.logger.Error(err)
		return fmt.Errorf("Password Encryption  failed")
	}
	account.Password = string(pass)

	if dbc := p.db.Debug().Create(&account); dbc.Error != nil {
		p.logger.Error(dbc.Error)
		return dbc.Error
	}
	role := entity.Role{}
	p.db.Debug().First(&role, roleID)

	permissions := []entity.Permission{}
	p.db.Debug().Model(&role).Related(&permissions)

	return nil
}
