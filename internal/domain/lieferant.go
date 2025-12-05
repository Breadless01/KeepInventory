package domain

type Lieferant struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
	Sitz string `json:"Sitz"`
}

type LieferantFilterResult struct {
	Items  []*Lieferant             `json:"items"`
	Total  int                      `json:"total"`
	Facets map[string][]FacetOption `json:"facets"`
}
