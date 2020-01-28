package repository

import (
	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/entity"
)

// Account accesses users' information from storage, such as database.
type Account interface {
	GetAcccount() string
	CreateAccount(req *proto.PostAccountRequest) (entity.Account, error)
}
