package main

import (
	"context"
	"embed"
	"log"
	"os"
	"path/filepath"

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
	exePath, err := os.Executable()
	if err != nil {
		exePath, _ = os.Getwd()
	}
	baseDir := filepath.Dir(exePath)
	db := sqliteadapter.OpenDB(filepath.Join(baseDir, "inventory.db"))

	searchRepo := sqliteadapter.NewSearchRepositorySQLite(db)
	lieferantRepo := sqliteadapter.NewLieferantRepositorySQLite(db)
	lieferantBauteilRepo := sqliteadapter.NewLieferantBauteilRepositorySQLite(db)
	typRepo := sqliteadapter.NewTypRepositorySQLite(db)
	artRepo := sqliteadapter.NewHerstellungsartRepositorySQLite(db)
	verschRepo := sqliteadapter.NewVerschleissteilRepositorySQLite(db)
	funktionRepo := sqliteadapter.NewFunktionRepositorySQLite(db)
	materialRepo := sqliteadapter.NewMaterialRepositorySQLite(db)
	oberfRepo := sqliteadapter.NewOberflaechenbehandlungRepositorySQLite(db)
	farbeRepo := sqliteadapter.NewFarbeRepositorySQLite(db)
	reserveRepo := sqliteadapter.NewReserveRepositorySQLite(db)

	bauteilRepo := sqliteadapter.NewBauteilRepositorySQLite(db)

	searchService := application.NewSearchService(searchRepo)
	bauteilService := application.NewBauteilService(
		bauteilRepo,
		lieferantBauteilRepo,
		typRepo,
		artRepo,
		verschRepo,
		funktionRepo,
		materialRepo,
		oberfRepo,
		farbeRepo,
		reserveRepo,
	)

	filterCfgService := application.NewFilterConfigService(db, baseDir)

	lieferantService := application.NewLieferantService(lieferantRepo)

	kundeRepo := sqliteadapter.NewKundeRepositorySQLite(db)
	kundeService := application.NewKundeService(kundeRepo)

	projektRepo := sqliteadapter.NewProjektRepositorySQLite(db)
	projektService := application.NewProjektService(projektRepo, kundeRepo)

	// Stammdaten services
	typService := application.NewTypService(typRepo)
	herstellungsartService := application.NewHerstellungsartService(artRepo)
	verschleissteilService := application.NewVerschleissteilService(verschRepo)
	funktionService := application.NewFunktionService(funktionRepo)
	materialService := application.NewMaterialService(materialRepo)
	oberflaechenbehandlungService := application.NewOberflaechenbehandlungService(oberfRepo)
	farbeService := application.NewFarbeService(farbeRepo)
	reserveService := application.NewReserveService(reserveRepo)

	backendApp := &backend.App{
		SearchService:                 searchService,
		BauteilService:                bauteilService,
		LieferantService:              lieferantService,
		KundeService:                  kundeService,
		ProjektService:                projektService,
		TypService:                    typService,
		HerstellungsartService:        herstellungsartService,
		VerschleissteilService:        verschleissteilService,
		FunktionService:               funktionService,
		MaterialService:               materialService,
		OberflaechenbehandlungService: oberflaechenbehandlungService,
		FarbeService:                  farbeService,
		ReserveService:                reserveService,
		FilterConfigService:           filterCfgService,
	}

	return &AppContainer{
		BackendApp: backendApp,
	}
}

func (a *AppContainer) Startup(ctx context.Context) {
	log.Println("AppContainer.Startup")
	a.BackendApp.Startup(ctx)
}
