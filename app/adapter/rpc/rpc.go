package rpc

import (
	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/elf/fw"
	"google.golang.org/grpc"
)

var _ fw.GRpcAPI = (*AccountAPI)(nil)

type AccountAPI struct {
	accountServer proto.AccountServer
}

func (a AccountAPI) RegisterServers(server *grpc.Server) {
	proto.RegisterAccountServer(server, a.accountServer)
}

func NewAccountAPI(accountServer proto.AccountServer) AccountAPI {
	return AccountAPI{accountServer: accountServer}
}
