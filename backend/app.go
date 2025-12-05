package backend

import (
	"context"
	"log"

	"KeepInventory/internal/application"
	"KeepInventory/internal/domain"
)

type App struct {
	ctx                           context.Context
	SearchService                 *application.SearchService
	BauteilService                *application.BauteilService
	LieferantBauteilService       *application.LieferantBauteilService
	KundeService                  *application.KundeService
	ProjektService                *application.ProjektService
	LieferantService              *application.LieferantService
	TypService                    *application.TypService
	HerstellungsartService        *application.HerstellungsartService
	VerschleissteilService        *application.VerschleissteilService
	FunktionService               *application.FunktionService
	MaterialService               *application.MaterialService
	OberflaechenbehandlungService *application.OberflaechenbehandlungService
	FarbeService                  *application.FarbeService
	ReserveService                *application.ReserveService
	FilterConfigService           *application.FilterConfigService
}

func (a *App) Startup(ctx context.Context) {
	log.Println("App starting up...")
	a.ctx = ctx
}

// SearchEngine
func (a *App) Search(query, objectType string, limit int) ([]domain.SearchResult, error) {
	req := domain.SearchRequest{
		Query:      query,
		ObjectType: objectType, // "" = alle
		Limit:      limit,
	}
	return a.SearchService.Search(req)
}

// View Models für Requests/Responses
type CreateBauteilRequest struct {
	ID             int64   `json:"ID"`
	TeilName       string  `json:"TeilName"`
	KundeId        int64   `json:"KundeId"`
	ProjektId      int64   `json:"ProjektId"`
	LieferantenIds []int64 `json:"LieferantenIds"`

	TypID                    int64 `json:"TypID"`
	HerstellungsartID        int64 `json:"HerstellungsartID"`
	VerschleissteilID        int64 `json:"VerschleissteilID"`
	FunktionID               int64 `json:"FunktionID"`
	MaterialID               int64 `json:"MaterialID"`
	OberflaechenbehandlungID int64 `json:"OberflaechenbehandlungID"`
	FarbeID                  int64 `json:"FarbeID"`
	ReserveID                int64 `json:"ReserveID"`
}

//Bauteile

func (a *App) CreateBauteil(req CreateBauteilRequest) (*domain.Bauteil, error) {
	return a.BauteilService.CreateBauteil(application.CreateBauteilInput{
		TeilName:                 req.TeilName,
		KundeId:                  req.KundeId,
		ProjektId:                req.ProjektId,
		LieferantenIds:           req.LieferantenIds,
		TypID:                    req.TypID,
		HerstellungsartID:        req.HerstellungsartID,
		VerschleissteilID:        req.VerschleissteilID,
		FunktionID:               req.FunktionID,
		MaterialID:               req.MaterialID,
		OberflaechenbehandlungID: req.OberflaechenbehandlungID,
		FarbeID:                  req.FarbeID,
		ReserveID:                req.ReserveID,
	})
}

func (a *App) UpdateBauteil(req CreateBauteilRequest) (*domain.Bauteil, error) {
	return a.BauteilService.UpdateBauteil(application.CreateBauteilInput{
		ID:        req.ID,
		TeilName:  req.TeilName,
		KundeId:   req.KundeId,
		ProjektId: req.ProjektId,
	})
}

func (a *App) FilterBauteile(state domain.FilterState) (domain.BauteilFilterResult, error) {
	return a.BauteilService.FacetFilter(state)
}

// Kunden
type CreateKundeRequest struct {
	Name string `json:"name"`
	Sitz string `json:"sitz"`
}

func (a *App) CreateKunde(req CreateKundeRequest) (*domain.Kunde, error) {
	return a.KundeService.CreateKunde(
		req.Name,
		req.Sitz,
	)
}

func (a *App) ListKunden() ([]*domain.Kunde, error) {
	return a.KundeService.ListKunden()
}

func (a *App) FilterKunden(state domain.FilterState) (domain.KundeFilterResult, error) {
	return a.KundeService.FacetFilter(state)
}

//Projekte

type CreateProjektRequest struct {
	Name  string `json:"name"`
	Kunde string `json:"kunde"`
}

func (a *App) CreateProjekt(req CreateProjektRequest) (*domain.Projekt, error) {
	return a.ProjektService.CreateProjekt(
		req.Name,
		req.Kunde,
	)
}

func (a *App) ListProjekte() ([]*domain.Projekt, error) {
	return a.ProjektService.ListProjekte()
}

func (a *App) FilterProjekte(state domain.FilterState) (domain.ProjektFilterResult, error) {
	return a.ProjektService.FacetFilter(state)
}

func (a *App) ListTypen() ([]*domain.Typ, error) {
	return a.TypService.FindAll()
}

func (a *App) ListHerstellungsarten() ([]*domain.Herstellungsart, error) {
	return a.HerstellungsartService.FindAll()
}

func (a *App) ListVerschleissteile() ([]*domain.Verschleissteil, error) {
	return a.VerschleissteilService.FindAll()
}

func (a *App) ListFunktionen() ([]*domain.Funktion, error) {
	return a.FunktionService.FindAll()
}

func (a *App) ListMaterialien() ([]*domain.Material, error) {
	return a.MaterialService.FindAll()
}

func (a *App) ListOberflaechenbehandlungen() ([]*domain.Oberflaechenbehandlung, error) {
	return a.OberflaechenbehandlungService.FindAll()
}

func (a *App) ListFarben() ([]*domain.Farbe, error) {
	return a.FarbeService.FindAll()
}

func (a *App) ListReserven() ([]*domain.Reserve, error) {
	return a.ReserveService.FindAll()
}

func (a *App) GetFilterConfig() (domain.FilterConfig, error) {
	return a.FilterConfigService.Load()
}

func (a *App) SaveFilterConfig(cfg domain.FilterConfig) error {
	return a.FilterConfigService.Save(cfg)
}

type LieferantRequest struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
	Sitz string `json:"Sitz"`
}

func (a *App) CreateLieferant(req LieferantRequest) (*domain.Lieferant, error) {
	return a.LieferantService.CreateLieferant(
		req.Name,
		req.Sitz,
	)
}

func (a *App) UpdateLieferant(req LieferantRequest) (*domain.Lieferant, error) {
	return a.LieferantService.UpdateLieferant(application.LieferantInput{
		ID:   req.ID,
		Name: req.Name,
		Sitz: req.Sitz,
	})
}

func (a *App) ListLieferanten() ([]*domain.Lieferant, error) {
	return a.LieferantService.ListLieferanten()
}

func (a *App) FilterLieferanten(state domain.FilterState) (domain.LieferantFilterResult, error) {
	return a.LieferantService.FacetFilter(state)
}
