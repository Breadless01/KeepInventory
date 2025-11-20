package domain

import "time"

// Bauteil repräsentiert ein einzelnes Inventar-Teil in deinem Lager.
type Bauteil struct {
	ID           int64     `json:"ID"`
	TeilName     string    `json:"TeilName"`
	KundeID      int64     `json:"KundeID"`
	ProjektID    int64     `json:"ProjektID"`
	Erstelldatum time.Time `json:"Erstelldatum"`

	TypID                    int64 `json:"TypID"`
	HerstellungsartID        int64 `json:"HerstellungsartID"`
	VerschleissteilID        int64 `json:"VerschleissteilID"`
	FunktionID               int64 `json:"FunktionID"`
	MaterialID               int64 `json:"MaterialID"`
	OberflaechenbehandlungID int64 `json:"OberflaechenbehandlungID"`
	FarbeID                  int64 `json:"FarbeID"`
	ReserveID                int64 `json:"ReserveID"`

	Sachnummer string `json:"Sachnummer"`
}
