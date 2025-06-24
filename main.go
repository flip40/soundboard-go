package main

import (
	"embed"
	"log"

	application "github.com/flip40/soundboard-go/backend/app"
	"github.com/flip40/soundboard-go/backend/keycodes"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := application.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Soundboard",
		Width:  720,
		Height: 570,
		// MinWidth:          720,
		// MinHeight:         570,
		// MaxWidth:          1280,
		// MaxHeight:         740,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 0, G: 56, B: 71, A: 255},
		Assets:            assets,
		// LogLevel:          logger.DEBUG,
		OnStartup: app.Startup,
		// OnDomReady: app.domReady,
		OnShutdown: app.Shutdown,
		Bind: []interface{}{
			application.Exported(app),
			&keycodes.KeycodeHelper{},
			&keycodes.Keycode{},
		},
		EnumBind: []interface{}{
			keycodes.KeycodeGroups,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
