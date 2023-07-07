package operation

import (
	"github.com/ancogamer/app/structs/investments"
	"github.com/ancogamer/app/structs/payment"
)

type Income struct {
	Investments investments.Investments `json:"investmets"`
	Payment     payment.Payment         `json:"payment"`
}
