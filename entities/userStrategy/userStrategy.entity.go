package userstrategy

type UserStrategy struct {
	UserStrategy uint64 `json:"user"`
	TradeConfig  uint64 `json:"trade_config"`
	Enabled      bool   `json:"enabled"`
}
