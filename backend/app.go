package backend

import (
	"context"
	"log"

	"KeepInventory/internal/application"
	"KeepInventory/internal/domain"
)

type App struct {
	ctx            context.Context
	BauteilService *application.BauteilService
	KundeService   *application.KundeService
	ProjektService *application.ProjektService
}

// Wails ruft das beim Start auf und gibt dir den Context.
func (a *App) Startup(ctx context.Context) {
	log.Println("App starting up...")
	a.ctx = ctx
}

// View Models für Requests/Responses
type CreateBauteilRequest struct {
	TeilName  string `json:"TeilName"`
	KundeID   int64  `json:"KundeID"`
	ProjektID int64  `json:"ProjektID"`

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
		KundeID:                  req.KundeID,
		ProjektID:                req.ProjektID,
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
	Name     string `json:"name"`
	KundenID int64  `json:"kunden_id"`
}

func (a *App) CreateProjekt(req CreateProjektRequest) (*domain.Projekt, error) {
	return a.ProjektService.CreateProjekt(
		req.Name,
		req.KundenID,
	)
}

func (a *App) ListProjekte() ([]*domain.Projekt, error) {
	return a.ProjektService.ListProjekte()
}
