package rpc

import (
	"context"
	"fmt"

	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/app/usecase"
	"github.com/sentadmedia/elf/fw"
)

var _ proto.AccountServer = (*AccountServer)(nil)

// AccountServer Data type
type AccountServer struct {
	useCase usecase.UseCase
	logger  fw.Logger
}

// NewAccountServer Creates a new AccountServer instance
func NewAccountServer(
	useCase usecase.UseCase,
	logger fw.Logger,
) AccountServer {
	return AccountServer{
		useCase: useCase,
		logger:  logger,
	}
}

// RegisterAccount Create a new account
func (a AccountServer) RegisterAccount(ctx context.Context, req *proto.RegisterAccountRequest) (*proto.RegisterAccountResponse, error) {
	account := entity.Account{
		Username:    req.UserName,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsSuperuser: req.IsSuperUser,
	}

	a.logger.Debug(fmt.Sprintf("Account before insert=%+v", account))

	err := a.useCase.RegisterAccount(&account, req.RoleId)
	if err != nil {
		a.logger.Error(err)
		return &proto.RegisterAccountResponse{}, err
	}

	a.logger.Debug(fmt.Sprintf("Account before insert=%+v", account))

	return &proto.RegisterAccountResponse{}, nil
}
