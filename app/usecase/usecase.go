package usecase

import (
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/accounts"
	"github.com/sentadmedia/elf/fw"
)

// UseCase Data type
type UseCase struct {
	logger   fw.Logger
	producer accounts.Producer
	consumer accounts.Consumer
}

// NewUseCase creates a new UseCase instance
func NewUseCase(
	logger fw.Logger,
	producer accounts.Producer,
	consumer accounts.Consumer,
) UseCase {
	return UseCase{
		logger:   logger,
		producer: producer,
		consumer: consumer,
	}
}

// RegisterAccount create a new account using producer
func (u UseCase) RegisterAccount(username, email, password string) (entity.Account, error) {
	return u.producer.RegisterAccount(username, email, password)
}
