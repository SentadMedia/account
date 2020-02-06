package rpc

import (
	"context"

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

// SignIn Log in a user
func (a AccountServer) SignIn(ctx context.Context, req *proto.SignInRequest) (res *proto.SignInResponse, err error) {
	a.logger.Debugf("SignIn gRPC request req=%v", req)
	res.UserId, err = a.useCase.SignIn(req.UserName, req.Password)
	if err != nil {
		a.logger.Error(err)
	}
	return
}

// RegisterAccount Create a new account
func (a AccountServer) RegisterAccount(ctx context.Context, req *proto.RegisterAccountRequest) (response *proto.Void, err error) {
	account := entity.Account{
		Username:    req.UserName,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsSuperuser: req.IsSuperUser,
	}

	err = a.useCase.RegisterAccount(&account, req.RoleId)
	if err != nil {
		a.logger.Error(err)
	}

	return
}
