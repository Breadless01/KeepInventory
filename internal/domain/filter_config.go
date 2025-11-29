package domain

type FieldFilterConfig struct {
	Field   string `json:"field"`
	Label   string `json:"label"`
	Enabled bool   `json:"enabled"`
}

type ResourceFilterConfig struct {
	Resource FilterResource      `json:"resource"`
	Table    string              `json:"table"`
	Fields   []FieldFilterConfig `json:"fields"`
}

type FilterConfig struct {
	Resources []ResourceFilterConfig `json:"resources"`
}
