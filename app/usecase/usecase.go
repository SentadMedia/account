package usecase

import (
	"fmt"

	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase/accounts"
	"github.com/sentadmedia/elf/fw"
)

type UseCase struct {
	logger   fw.Logger
	producer accounts.Producer
	consumer accounts.Consumer
}

func (u UseCase) PostAccount(req *proto.PostAccountRequest) (entity.Account, error) {
	resp, err := u.producer.CreateAccount(req)
	u.logger.Debug(fmt.Sprintf("Creating a user res:%v err:%s", resp, err))
	return resp, nil
}

func (u UseCase) GetAccount() string {
	r, _ := u.consumer.ConsumeInBatch()
	u.logger.Debug(fmt.Sprintf("Response from DB: %s", r))
	return r
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
