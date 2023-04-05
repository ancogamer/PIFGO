package outcome

import (
	"github.com/ancogamer/app/structs/tax"
)

type Tax struct {
	Name string `json:"name"`
	tax.Tax
}
