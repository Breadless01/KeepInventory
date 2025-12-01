package domain

type Kunde struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
	Sitz string `json:"Sitz"`
}

type KundeFilterResult struct {
	Items  []*Kunde                 `json:"items"`
	Total  int                      `json:"total"`
	Facets map[string][]FacetOption `json:"facets"`
}
