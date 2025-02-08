package repositorydto

type InputUserStrategy struct {
	User        uint64 `json:"user"`
	TradeConfig uint64 `json:"trade_config"`
}

type OutputUserStrategy struct {
	UserStrategy uint64 `json:"user_strategy"`
	User         uint64 `json:"user"`
	TradeConfig  uint64 `json:"trade_config"`
}
