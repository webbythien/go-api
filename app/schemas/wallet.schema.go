package schemas

type BalanceResponse struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}
