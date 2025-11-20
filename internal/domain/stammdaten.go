// internal/domain/stammdaten.go
package domain

type Typ struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Herstellungsart struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Verschleissteil struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Funktion struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Material struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Oberflaechenbehandlung struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Farbe struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}

type Reserve struct {
	ID     int64  `json:"ID"`
	Name   string `json:"Name"`
	Symbol int    `json:"Symbol"`
}
