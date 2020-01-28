package repository

// Account accesses users' information from storage, such as database.
type Account interface {
	GetAcccount() string
}
