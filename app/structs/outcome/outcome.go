package outcome

type OutCome struct {
	Card    *Card    `json:"card,omitempty"`
	Generic *Generic `json:"generic,omitempty"`
	Tax     *Tax     `json:"tax,omitempty"`
	Value   float64  `json:"value" example:"1.00"`
}
