package domain

type Projekt struct {
	ID      int64  `json:"ID"`
	Name    string `json:"Name"`
	KundeID int64  `json:"KundeID"`

	Kunde *Kunde `json:"Kunde,omitempty"`
}
