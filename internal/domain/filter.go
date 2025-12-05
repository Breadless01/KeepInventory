package domain

type FilterResource string

const (
	ResourceBauteile    FilterResource = "bauteile"
	ResourceKunden      FilterResource = "kunden"
	ResourceProjekte    FilterResource = "projekte"
	ResourceLieferanten FilterResource = "lieferanten"
)

// Zustand der Filter im UI
type FilterState struct {
	Resource FilterResource `json:"resource"` // z.B. "bauteile"
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
	// Feldname -> Liste von IDs (Multi-Select)
	Filters map[string][]any `json:"filters"`
}

// Eine Facet-Option: z.B. Material = „Edelstahl (5)“
type FacetOption struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
