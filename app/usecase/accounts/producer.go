package accounts

import (
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
	"github.com/sentadmedia/elf/fw"
)

// Producer interface
type Producer interface {
	RegisterAccount(account *entity.Account, roleID int32) error
	SignIn(username, password string) error
}

var _ Producer = (*ProducerPersist)(nil)

// ProducerPersist interface
type ProducerPersist struct {
	logger fw.Logger
	repo   repository.Account
}

// RegisterAccount write account into DB
func (p ProducerPersist) RegisterAccount(account *entity.Account, roleID int32) error {
	return p.repo.RegisterAccount(account, roleID)
}

// SignIn Log In a user
func (p ProducerPersist) SignIn(username, password string) error {
	return p.repo.SignIn(username, password)
}

// NewProducerPersist creates and initializes Producer
func NewProducerPersist(
	repo repository.Account,
	logger fw.Logger,
) ProducerPersist {
	return ProducerPersist{
		repo:   repo,
		logger: logger,
	}
}
