package model

type Preduzece struct {
}
type Ponuda struct {
	Cena              float64   `json:"cena"`
	Ponudjac          Preduzece `json:"ponudjac"`
	RokINacinPlacanja string    `json:"rok_i_nacin_placanja"`
}
type PlanJavneNabavke struct {
	TipPredmeta        string `json:"tip_predmeta"`
	ProcenjenaVrednost string `json:"procenjena_vrednost"`
	Kolicina           int    `json:"kolicina"`
}

type Nabavka struct {
	Potrazilac          Preduzece        `json:"potrazilac"`
	Ponude              []Ponuda         `json:"potrazilac"`
	PredmetNabavke      PlanJavneNabavke `json:"predmet_nabavke"`
	datumPocetkaNabavke string           `json:"datum_pocetka_nabavke"`
	datumKrajaNabavke   string           `json:"datum_kraja_nabavke"`
	nazivNabavke        string           `json:"naziv_nabavke"`
	kratakOpis          string           `json:"kratak_opis"`
}
