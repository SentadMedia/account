package accounts

import (
	"github.com/sentadmedia/account/app/usecase/repository"
)

type Consumer interface {
}

var _ Consumer = (*ConsumerPersist)(nil)

type ConsumerPersist struct {
	repo repository.Account
}

func NewConsumerPersist(
	repo repository.Account,
) ConsumerPersist {
	return ConsumerPersist{
		repo: repo,
	}
}
