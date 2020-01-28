package accounts

import (
	"github.com/sentadmedia/account/app/usecase/repository"
	"github.com/sentadmedia/elf/fw"
)

type Producer interface {
	Produce(KeyLength uint) error
}

var _ Producer = (*ProducerPersist)(nil)

// ProducerPersist represents persistent key producer
type ProducerPersist struct {
	logger fw.Logger
	repo   repository.Account
}

// Produce generates unique keys and store them in the repository
func (p ProducerPersist) Produce(KeyLength uint) error {
	return nil
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
