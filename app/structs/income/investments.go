package income

import (
	"github.com/ancogamer/app/structs/tax"
)

type FixReturn struct {
	Name     string  `json:"name" example:"Treasure"`
	TaxYear  tax.Tax `json:"tax_year" example:"13.65"`
	TaxMonth tax.Tax `json:"tax_month" example:"1.2"`
}

// TODO support variable income, like stocks
//type VariableReturn struct {
//	Name string  `json:"name" example:"Treasure"`
//	Min  float64 `json:"min" example:"1.0"`
//	Max  float64 `json:"max" example:"5.0"`
//}
