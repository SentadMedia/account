package accounts

import (
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
	"github.com/sentadmedia/elf/fw"
)

// Producer interface
type Producer interface {
	RegisterAccount(username, email, password string) (entity.Account, error)
}

var _ Producer = (*ProducerPersist)(nil)

// ProducerPersist interface
type ProducerPersist struct {
	logger fw.Logger
	repo   repository.Account
}

// RegisterAccount write account into DB
func (p ProducerPersist) RegisterAccount(username, email, password string) (entity.Account, error) {
	return p.repo.RegisterAccount(username, email, password)
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
