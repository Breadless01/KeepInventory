package application

import "KeepInventory/internal/domain"

type SearchRepository interface {
	Search(req domain.SearchRequest) ([]domain.SearchResult, error)
}

type BauteilRepository interface {
	Create(bauteil *domain.Bauteil) (*domain.Bauteil, error)
	Update(bauteil *domain.Bauteil) (*domain.Bauteil, error)
	CountByAttributes(
		typID, herstellungsartID, verschleissteilID,
		funktionID, materialID, oberflaechenbehandlungID,
		farbeID, reserveID int64,
	) (int64, error)
	FindByFilter(req domain.FilterState) ([]*domain.Bauteil, error)
	GetAttributeValuesById(facets map[string]map[int64]int) map[string]map[int64]string
}

type LieferantRepository interface {
	Create(lieferant *domain.Lieferant) (*domain.Lieferant, error)
	Update(lieferant *domain.Lieferant) (*domain.Lieferant, error)
	FindAll() ([]*domain.Lieferant, error)
	FindById(id int64) (*domain.Lieferant, error)
	FindByFilter(req domain.FilterState) ([]*domain.Lieferant, error)
}

type LieferantBauteilRepository interface {
	Create(lieferantBauteil *domain.LieferantBauteil) error
	Delete(bauteilId int64, lieferantId int64) error
	FindByBauteilId(id int64) ([]*domain.LieferantBauteil, error)
	FindByLieferantId(id int64) ([]*domain.LieferantBauteil, error)
}

type TypRepository interface {
	FindByID(id int64) (*domain.Typ, error)
	FindAll() ([]*domain.Typ, error)
}

type HerstellungsartRepository interface {
	FindByID(id int64) (*domain.Herstellungsart, error)
	FindAll() ([]*domain.Herstellungsart, error)
}

type VerschleissteilRepository interface {
	FindByID(id int64) (*domain.Verschleissteil, error)
	FindAll() ([]*domain.Verschleissteil, error)
}

type FunktionRepository interface {
	FindByID(id int64) (*domain.Funktion, error)
	FindAll() ([]*domain.Funktion, error)
}

type MaterialRepository interface {
	FindByID(id int64) (*domain.Material, error)
	FindAll() ([]*domain.Material, error)
}

type OberflaechenbehandlungRepository interface {
	FindByID(id int64) (*domain.Oberflaechenbehandlung, error)
	FindAll() ([]*domain.Oberflaechenbehandlung, error)
}

type FarbeRepository interface {
	FindByID(id int64) (*domain.Farbe, error)
	FindAll() ([]*domain.Farbe, error)
}

type ReserveRepository interface {
	FindByID(id int64) (*domain.Reserve, error)
	FindAll() ([]*domain.Reserve, error)
}

type KundeRepository interface {
	Create(k *domain.Kunde) (*domain.Kunde, error)
	FindByID(id int64) (*domain.Kunde, error)
	FindAll() ([]*domain.Kunde, error)
	FindByFilter(req domain.FilterState) ([]*domain.Kunde, error)
}

type ProjektRepository interface {
	Create(p *domain.Projekt) (*domain.Projekt, error)
	FindByID(id int64) (*domain.Projekt, error)
	FindAll() ([]*domain.Projekt, error)
	FindByFilter(req domain.FilterState) ([]*domain.Projekt, error)
}
