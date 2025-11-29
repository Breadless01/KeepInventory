package backend

import (
	"context"
	"log"

	"KeepInventory/internal/application"
	"KeepInventory/internal/domain"
)

type App struct {
	ctx                           context.Context
	BauteilService                *application.BauteilService
	KundeService                  *application.KundeService
	ProjektService                *application.ProjektService
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

// Wails ruft das beim Start auf und gibt dir den Context.
func (a *App) Startup(ctx context.Context) {
	log.Println("App starting up...")
	a.ctx = ctx
}

// View Models für Requests/Responses
type CreateBauteilRequest struct {
	TeilName  string `json:"TeilName"`
	KundeId   int64  `json:"KundeId"`
	ProjektId int64  `json:"ProjektId"`

	TypID                    int64 `json:"TypID"`
	HerstellungsartID        int64 `json:"HerstellungsartID"`
	VerschleissteilID        int64 `json:"VerschleissteilID"`
	FunktionID               int64 `json:"FunktionID"`
	MaterialID               int64 `json:"MaterialID"`
	OberflaechenbehandlungID int64 `json:"OberflaechenbehandlungID"`
	FarbeID                  int64 `json:"FarbeID"`
	ReserveID                int64 `json:"ReserveID"`
}

func (a *App) CreateBauteil(req CreateBauteilRequest) (*domain.Bauteil, error) {
	return a.BauteilService.CreateBauteil(application.CreateBauteilInput{
		TeilName:                 req.TeilName,
		KundeId:                  req.KundeId,
		ProjektId:                req.ProjektId,
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

func (a *App) ListBauteile() ([]*domain.Bauteil, error) {
	return a.BauteilService.ListBauteile()
}

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
