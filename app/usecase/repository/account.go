package repository

import "github.com/sentadmedia/account/app/entity"

// Account accesses users' information from storage, such as database.
type Account interface {
	RegisterAccount(account *entity.Account, roleID int32) error
	SignIn(username, password string) error
}
