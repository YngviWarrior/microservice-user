package usecasesdto

type InputCreateUser struct {
	Name   string
	Email  string
	Active bool
}

type OutputCreateUser struct {
	User   uint64
	Name   string
	Email  string
	Active bool
}
