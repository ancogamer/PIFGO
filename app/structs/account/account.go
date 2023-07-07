package account

import (
	"github.com/ancogamer/app/structs/operation"
)

type Account struct {
	Name        string            `json:"name"`
	BalanceTemp float64           `json:"temp_balance"` // balance taking note of splitted ops (example: bought sometime but will pay the full price in 10 months(1/10 per month))
	RealBalance float64           `json:"real_balance"` // real money on the account
	Income      operation.Income  `json:"income"`
	OutCome     operation.OutCome `json:"outcome"`
}
