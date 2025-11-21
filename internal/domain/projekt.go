package domain

type Projekt struct {
	ID    int64  `json:"ID"`
	Name  string `json:"Name"`
	Kunde string `json:"Kunde"`
}
