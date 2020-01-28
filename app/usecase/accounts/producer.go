package accounts

import (
	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/repository"
	"github.com/sentadmedia/elf/fw"
)

type Producer interface {
	CreateAccount(req *proto.PostAccountRequest) (entity.Account, error)
}

var _ Producer = (*ProducerPersist)(nil)

// ProducerPersist represents persistent key producer
type ProducerPersist struct {
	logger fw.Logger
	repo   repository.Account
}

// Produce generates unique keys and store them in the repository
func (p ProducerPersist) CreateAccount(req *proto.PostAccountRequest) (entity.Account, error) {
	return p.repo.CreateAccount(req)
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
