package domain

type LieferantBauteil struct {
	ID         int64  `json:"ID"`
	LiferantId int64  `json:"LiferantId"`
	Lieferant  string `json:"Lieferant"`
	BauteilId  int64  `json:"BauteilId"`
	Bauteil    string `json:"Bauteil"`
}

type LieferantBauteilFilterResult struct {
	Items  []*LieferantBauteil      `json:"items"`
	Total  int                      `json:"total"`
	Facets map[string][]FacetOption `json:"facets"`
}
