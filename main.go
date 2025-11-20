package main

import (
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func main() {
	log.Println("KeepInventory Wails main startet...")

	appContainer := NewAppContainer()

	err := wails.Run(&options.App{
		Title:  "KeepInventory",
		Width:  1024,
		Height: 768,
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: appContainer.Startup,
		Bind: []interface{}{
			appContainer.BackendApp,
		},
	})

	if err != nil {
		log.Fatalf("Wails Run Fehler: %v", err)
	}
}
