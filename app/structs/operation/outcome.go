package operation

import (
	"github.com/ancogamer/app/structs/card"
	"github.com/ancogamer/app/structs/generic"
	"github.com/ancogamer/app/structs/tax"
)

type OutCome struct {
	CardBills []card.CardBill   `json:"cardbills"`
	Taxes     []tax.Tax         `json:"taxes"`
	Generic   []generic.Generic `json:"Generic"`
}
