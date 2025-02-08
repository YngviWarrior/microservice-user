package controllerdto

type InputCreateUser struct {
	Name   *string `json:"name" validate:"required"`
	Email  *string `json:"email" validate:"required,email"`
	Active *bool   `json:"active" validate:""`
}

type OutCreateUser struct{}
