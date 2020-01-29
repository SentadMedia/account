package rpc

import (
	"context"
	"fmt"

	"github.com/sentadmedia/account/app/adapter/rpc/proto"
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
	account, err := a.useCase.RegisterAccount(req.Username, req.Password, req.Email)
	if err != nil {
		a.logger.Error(fmt.Errorf("Error whire creating an account err=%s", err.Error()))
		return &proto.RegisterAccountResponse{}, err
	}

	return &proto.RegisterAccountResponse{
		Account: &proto.User{
			Id:       account.ID,
			Username: account.Username,
			Email:    account.Email,
			Password: account.Password,
		},
	}, nil
}
