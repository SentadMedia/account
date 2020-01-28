package accounts

import (
	"github.com/sentadmedia/account/app/usecase/repository"
)

type Consumer interface {
	ConsumeInBatch() (string, error)
}

var _ Consumer = (*ConsumerPersist)(nil)

type ConsumerPersist struct {
	repo repository.Account
}

func (p ConsumerPersist) ConsumeInBatch() (string, error) {
	return p.repo.GetAcccount(), nil
}

func NewConsumerPersist(
	repo repository.Account,
) ConsumerPersist {
	return ConsumerPersist{
		repo: repo,
	}
}
