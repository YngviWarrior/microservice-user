package repositorydto

type InputUser struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

type OutputUser struct {
	User   string `json:"user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}
