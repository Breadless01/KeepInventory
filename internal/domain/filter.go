package domain

// Welche Ressource wollen wir filtern?
type FilterResource string

const (
	ResourceBauteile FilterResource = "bauteile"
	ResourceKunden   FilterResource = "kunden"
	ResourceProjekte FilterResource = "projekte"
)

// Zustand der Filter im UI
type FilterState struct {
	Resource FilterResource `json:"resource"` // z.B. "bauteile"
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
	// Feldname -> Liste von IDs (Multi-Select)
	Filters map[string][]int64 `json:"filters"`
}

// Eine Facet-Option: z.B. Material = „Edelstahl (5)“
type FacetOption struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Ergebnis eines Facet-Filters
type FilterResult struct {
	Items  any                      `json:"items"`
	Total  int                      `json:"total"`
	Facets map[string][]FacetOption `json:"facets"`
}
