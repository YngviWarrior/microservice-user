package userstrategy

type UserStrategy struct {
	User     uint64 `json:"user"`
	Strategy uint64 `json:"strategy"`
	Enabled  bool   `json:"enabled"`
}
