package controllerdto

type InputGetUserByEmail struct {
	Email *string `json:"email" validate:"required,email"`
}

type OutputGetUserByEmail struct {
	User   uint64 `json:"user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}
