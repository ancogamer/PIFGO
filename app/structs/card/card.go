package card

import (
	"github.com/ancogamer/app/structs/tax"
)

type Card struct {
	Alias       string  `json:"alias" example:"Card wiht final 1234"` // Just a alias for the card
	GenMile     bool    `json:"gen_mile" example:"true"`              // Info if that card generates miles
	QtdMile     *Mile   // Representation of a Mile, will be nil case card dont generate miles
	Tax         tax.Tax `json:"tax"`          // most cards have fees to be used
	AllowCredit bool    `json:"allow_credit"` // used to make possible subdivision
}

type Mile struct {
	Points int8 `json:"points" example:"1"`  // Points that the card generates
	PerQTD int8 `json:"per_qtd" example:"3"` // QTD necessary to generate card points
}

type CardBill struct {
	Card `json:"card"`
	Ops  []CardOps
}

type CardOps struct {
	OpTDebit   bool    `json:"op_t_debit"`
	OpTCredit  bool    `json:"op_t_credit"`
	Value      float64 `json:"value"`
	QtdPayment int     `json:"qtd_payment"`
}
