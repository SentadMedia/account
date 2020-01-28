package usecase

import (
	"github.com/sentadmedia/account/app/usecase/accounts"
	"github.com/sentadmedia/elf/fw"
)

type UseCase struct {
	logger   fw.Logger
	producer accounts.Producer
	consumer accounts.Consumer
}

func (u UseCase) GetAccount() string {
	r, _ := u.consumer.ConsumeInBatch()
	u.logger.Info(r)
	return "mgh.soufiane@gmail.com"
}

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
