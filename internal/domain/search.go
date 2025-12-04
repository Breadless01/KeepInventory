package domain

type SearchRequest struct {
	Query      string
	ObjectType string
	Limit      int
}

type SearchResult struct {
	ID       int64             `json:"id"`
	Type     string            `json:"type"`
	Label    string            `json:"label"`
	Subtitle string            `json:"subtitle"`
	Extra    map[string]string `json:"extra"`
}
