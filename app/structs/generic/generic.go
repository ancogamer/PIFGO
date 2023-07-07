package generic

type Generic struct {
	Name    string  `json:"name" example:"name"`
	Value   float64 `json:"value" example:"10"`
	Destine string  `json:"destine" example:"Local store" `
}
