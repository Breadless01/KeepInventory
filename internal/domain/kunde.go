package domain

type Kunde struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
	Sitz string `json:"Sitz"`
}
