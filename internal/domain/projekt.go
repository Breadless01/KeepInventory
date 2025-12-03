package domain

type Projekt struct {
	ID    int64  `json:"ID"`
	Name  string `json:"Name"`
	Kunde string `json:"Kunde"`
}

type ProjektFilterResult struct {
	Items  []*Projekt               `json:"items"`
	Total  int                      `json:"total"`
	Facets map[string][]FacetOption `json:"facets"`
}
