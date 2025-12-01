package domain

type Bauteil struct {
	ID           int64  `json:"ID"`
	TeilName     string `json:"TeilName"`
	Kunde        string `json:"Kunde"`
	KundeId      int64  `json:"KundeId"`
	Projekt      string `json:"Projekt"`
	ProjektId    int64  `json:"ProjektId"`
	Erstelldatum string `json:"Erstelldatum"`

	TypID                    int64 `json:"TypID"`
	HerstellungsartID        int64 `json:"HerstellungsartID"`
	VerschleissteilID        int64 `json:"VerschleissteilID"`
	FunktionID               int64 `json:"FunktionID"`
	MaterialID               int64 `json:"MaterialID"`
	OberflaechenbehandlungID int64 `json:"OberflaechenbehandlungID"`
	FarbeID                  int64 `json:"FarbeID"`
	ReserveID                int64 `json:"ReserveID"`

	Sachnummer string `json:"Sachnummer"`
}

type BauteilFilterResult struct {
	Items  []*Bauteil               `json:"items"`
	Total  int                      `json:"total"`
	Facets map[string][]FacetOption `json:"facets"`
}
