package userEntity

type User struct {
	User   uint64 `json:"user"`
	Name   string `json:"name"`
	Active string `json:"active"`
}
