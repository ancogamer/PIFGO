package tax

type Tax struct {
	Name               string  `json:"name" example:"some name"`
	PorcentagePerYear  float64 `json:"por_per_year" example:"1.2"`
	PorcentagePerMonth float64 `json:"por_per_month" example:"1.2"`
	Value              float64 `json:"value" example:"10"`
}
