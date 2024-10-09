package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	configFileStorage, err := InitNewConfigFileStorage("MemoCraft")
    if err != nil {
        fmt.Printf("could not initialize the config store: %v\n", err)
        return
    }

	appMenu := createApplicationMenu(app)

	// Create application with options
	err = wails.Run(&options.App{
		Title:             "MemoCraft Notes",
		Menu:              appMenu,
		Width:             1024,
		Height:            768,
		WindowStartState:  options.Maximised,
		StartHidden:       false,
		HideWindowOnClose: false,
		AlwaysOnTop:       false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			configFileStorage,
		},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop:     false, // TODO: consider drag and drop functionality
			DisableWebViewDrop: false,
			CSSDropProperty:    "--wails-drop-target",
			CSSDropValue:       "drop",
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "MemoCraft Notes",
				Message: "© 2024 Dmytro Feschenko <feschenko.dmitryi@gmail.com>",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "MemoCraft",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
