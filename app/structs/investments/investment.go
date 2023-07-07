package investments

import "github.com/ancogamer/app/structs/tax"

type Investments struct {
	FixReturns []FixReturn `json:"fix_return"`
}

// ------------------------------------------------------------------------------------------------------------
type FixReturn struct {
	Name      string    `json:"name" example:"Treasure"`
	ReturnTax tax.Tax   `json:"return_tax"`
	Operation []tax.Tax `json:"tax"` // can more than 1 operation tax, like multiple government tax
}
