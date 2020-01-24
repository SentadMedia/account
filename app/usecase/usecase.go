package usecase

import (
	"github.com/sentadmedia/elf/fw"
	// "github.com/short-d/kgs/app/adapter/message"
	// "github.com/short-d/kgs/app/entity"
	// "github.com/short-d/kgs/app/usecase/keys"
	// "github.com/short-d/kgs/app/usecase/notification"
)

type UseCase struct {
	logger fw.Logger
}

func (u UseCase) GetAccount() string {
	u.logger.Info("Getting Account")
	return "mgh.soufiane@gmail.com"
}

func NewUseCase(
	logger fw.Logger,
) UseCase {
	return UseCase{
		logger: logger,
	}
}
