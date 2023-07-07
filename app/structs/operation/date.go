package operation

type Year struct {
	Year     string  `json:"year" example:"2023"` // Year in context. example 2023
	Months   []Month `json:"months"`
	TOutcome float64 `json:"total_outcome"`
	TIncome  float64 `json:"total_income"`
}

type Month struct {
	Name         string  `json:"name"`
	NumberInYear int     `json:"num_in_year"`
	Days         []Day   `json:"days"` // group of days in that month
	TOutcome     float64 `json:"total_outcome" example:"100.00"`
	TIncome      float64 `json:"total_income" example:"200.00"`
}
type Day struct {
	Name       string `json:"name" example:"monday"`     // representation in language of that day.
	NumInMonth int8   `json:"num_in_month" example:"01"` // number of the day in that year
	Income     []Income
	OutCome    []OutCome
}
