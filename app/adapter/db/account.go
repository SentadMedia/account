package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
	"github.com/sentadmedia/elf/fw"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// SignIn Check if the user exists and if it have the correct password, then creates a sessionToken
func (p AccountSQL) SignIn(username, password string) (uint64, error) {
	// Check if the user exists
	var user entity.Account
	if err := p.db.First(&user, "username = ?", username).Error; gorm.IsRecordNotFoundError(err) {
		p.logger.Errorf("Cannot find a user with the provided username (%s)", username)
		return 0, status.Error(codes.InvalidArgument, "username and/or password does not match")
	} else if err != nil {
		p.logger.Errorf("Unexpected error while querying for a username with value (%s) err=%v", username, err)
		return 0, status.Error(codes.Internal, "Internal server error")
	}

	// Verify it have the correct password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		p.logger.Warnf("Wrong password provided for user (%s)", username)
		return 0, status.Error(codes.InvalidArgument, "username and/or password does not match")
	}

	return uint64(user.ID), nil
}

// RegisterAccount Actually writes a user into DB
func (p AccountSQL) RegisterAccount(account *entity.Account, roleID int32) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.MinCost)
	if err != nil {
		p.logger.Error(err)
		return fmt.Errorf("Password Encryption  failed")
	}
	account.Password = string(pass)

	if err := p.db.Debug().Create(&account).Error; err != nil {
		return err
	}
	role := entity.Role{}
	if err := p.db.Debug().Find(&role, roleID).Error; gorm.IsRecordNotFoundError(err) {
		return fmt.Errorf(fmt.Sprintf("Unable to find a role using ID=%d", roleID))
	} else if err != nil {
		return err
	}

	permissions := []entity.Permission{}
	if err := p.db.Debug().Model(&role).Related(&permissions).Error; gorm.IsRecordNotFoundError(err) {
		return fmt.Errorf(fmt.Sprintf("There is no permissions attached to roleID=%d", roleID))
	} else if err != nil {
		return err
	}

	account.Permissions = permissions
	p.db.Debug().Save(&account)
	return nil
}
