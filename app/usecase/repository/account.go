package repository

import "github.com/sentadmedia/account/app/entity"

// Account accesses users' information from storage, such as database.
type Account interface {
	RegisterAccount(username, email, password string) (entity.Account, error)
}
