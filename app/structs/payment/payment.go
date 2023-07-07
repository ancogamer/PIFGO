package payment

import (
	"github.com/ancogamer/app/structs/generic"
	"github.com/ancogamer/app/structs/tax"
)

type Payment struct {
	Payments []PaymentOps `json:"payment_ops"`
}
type PaymentOps struct {
	generic.Generic
	Tax []tax.Tax `json:"taxes"`
}
