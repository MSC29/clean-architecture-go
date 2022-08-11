package cat_facts

type CatFactPresenter struct {
	Fact    string `json:"fact"`
	NbChars int32  `json:"nb_chars"`
}
