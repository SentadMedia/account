package rpc

import (
	"context"

	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/usecase"
	"github.com/sentadmedia/elf/fw"
	"google.golang.org/grpc"
)

var _ proto.AccountAPIServer = (*AccountAPIServer)(nil)

type AccountAPIServer struct {
	useCase usecase.UseCase
}

func (a AccountAPIServer) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	accountName := a.useCase.GetAccount()
	return &proto.GetAccountResponse{Account: accountName}, nil
}

func NewAccountAPIServer(useCase usecase.UseCase) AccountAPIServer {
	return AccountAPIServer{
		useCase: useCase,
	}
}

var _ fw.GRpcAPI = (*AccountAPI)(nil)

type AccountAPI struct {
	accountAPIServer proto.AccountAPIServer
}

func (a AccountAPI) RegisterServers(server *grpc.Server) {
	proto.RegisterAccountAPIServer(server, a.accountAPIServer)
}

func NewAccountAPI(accountAPIServer proto.AccountAPIServer) AccountAPI {
	return AccountAPI{accountAPIServer: accountAPIServer}
}
