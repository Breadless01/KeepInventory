package application

import "KeepInventory/internal/domain"

type BauteilRepository interface {
	Create(bauteil *domain.Bauteil) (*domain.Bauteil, error)
	FindAll() ([]*domain.Bauteil, error)
	CountByAttributes(
		typID, herstellungsartID, verschleissteilID,
		funktionID, materialID, oberflaechenbehandlungID,
		farbeID, reserveID int64,
	) (int64, error)
	FindByFilter(req domain.FilterState) ([]*domain.Bauteil, error)
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
}

type ProjektRepository interface {
	Create(p *domain.Projekt) (*domain.Projekt, error)
	FindByID(id int64) (*domain.Projekt, error)
	FindAll() ([]*domain.Projekt, error)
	FindByKunde(kunde string) ([]*domain.Projekt, error)
}
