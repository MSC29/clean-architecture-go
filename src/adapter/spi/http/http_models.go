package http

type CatFactsApiModel struct {
	CurrentPage int32 `json:"current_page"`
	Data        []CatFactApiModel
}

type CatFactApiModel struct {
	Fact   string `json:"fact"`
	Length int32  `json:"length"`
}
