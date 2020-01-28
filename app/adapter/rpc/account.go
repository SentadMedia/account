package rpc

import (
	"context"

	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/usecase"
)

var _ proto.AccountServer = (*AccountServer)(nil)

type AccountServer struct {
	useCase usecase.UseCase
}

func (a AccountServer) PostAccount(ctx context.Context, req *proto.PostAccountRequest) (*proto.PostAccountResponse, error) {
	account, err := a.useCase.PostAccount(req)
	return &proto.PostAccountResponse{
		Account: &proto.User{
			Id:       account.ID,
			Username: account.Username,
			Email:    account.Email,
			Password: account.Password,
		},
	}, err
}

func (a AccountServer) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	accountName := a.useCase.GetAccount()
	return &proto.GetAccountResponse{Account: accountName}, nil
}

func NewAccountServer(
	useCase usecase.UseCase,
) AccountServer {
	return AccountServer{
		useCase: useCase,
	}
}
