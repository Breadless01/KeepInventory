package main

import (
	"context"
	"embed"
	"log"

	"KeepInventory/backend"
	"KeepInventory/internal/application"
	sqliteadapter "KeepInventory/internal/infrastructure/sqlite"
)

//go:embed frontend/dist
var assets embed.FS

// AppContainer bindet Hex-Backend + Wails-Backend-App zusammen.
type AppContainer struct {
	BackendApp *backend.App
}

func NewAppContainer() *AppContainer {
	db := sqliteadapter.OpenDB("inventory.db")

	typRepo := sqliteadapter.NewTypRepositorySQLite(db)
	artRepo := sqliteadapter.NewHerstellungsartRepositorySQLite(db)
	verschRepo := sqliteadapter.NewVerschleissteilRepositorySQLite(db)
	funktionRepo := sqliteadapter.NewFunktionRepositorySQLite(db)
	materialRepo := sqliteadapter.NewMaterialRepositorySQLite(db)
	oberfRepo := sqliteadapter.NewOberflaechenbehandlungRepositorySQLite(db)
	farbeRepo := sqliteadapter.NewFarbeRepositorySQLite(db)
	reserveRepo := sqliteadapter.NewReserveRepositorySQLite(db)

	bauteilRepo := sqliteadapter.NewBauteilRepositorySQLite(db)

	bauteilService := application.NewBauteilService(
		bauteilRepo,
		typRepo,
		artRepo,
		verschRepo,
		funktionRepo,
		materialRepo,
		oberfRepo,
		farbeRepo,
		reserveRepo,
	)

	kundeRepo := sqliteadapter.NewKundeRepositorySQLite(db)
	kundeService := application.NewKundeService(kundeRepo)

	projektRepo := sqliteadapter.NewProjektRepositorySQLite(db)
	projektService := application.NewProjektService(projektRepo, kundeRepo)

	backendApp := &backend.App{
		BauteilService: bauteilService,
		KundeService:   kundeService,
		ProjektService: projektService,
	}

	return &AppContainer{
		BackendApp: backendApp,
	}
}

func (a *AppContainer) Startup(ctx context.Context) {
	log.Println("AppContainer.Startup")
	a.BackendApp.Startup(ctx)
}
