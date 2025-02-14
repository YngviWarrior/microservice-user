package usecasesdto

type InputGetUserByEmail struct {
	Email string
}

type OutputGetUserByEmail struct {
	User   uint64
	Name   string
	Email  string
	Active bool
}

type InputGetUserById struct {
	Id uint64
}

type OutputGetUserById struct {
	User   uint64
	Name   string
	Email  string
	Active bool
}
