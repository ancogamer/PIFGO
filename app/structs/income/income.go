package income

type Income struct {
	InvestmentsFix *FixReturn `json:"invest_fix,omitempty"`
	Value          float64    `json:"value" example:"1.00"`
}
