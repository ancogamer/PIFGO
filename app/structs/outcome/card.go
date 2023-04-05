package outcome

type Card struct {
	Alias   string `json:"alias" example:"Card wiht final 1234"` // Just a alias for the card
	GenMile bool   `json:"gen_mile" example:"true"`              // Info if that card generates miles
	QtdMile *Mile  // Representation of a Mile, will be nil case card dont generate miles
}

type Mile struct {
	Points int8 `json:"points" example:"1"`  // Points that the card generates
	PerQTD int8 `json:"per_qtd" example:"3"` // QTD necessary to generate card points
}
