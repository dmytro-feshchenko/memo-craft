package main

import (
	"embed"

	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	appMenu := createApplicationMenu(app)

	// Create application with options
	err := wails.Run(&options.App{
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

func createApplicationMenu(app *App) *menu.Menu {
	// Configure menu bar
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		wailsRuntime.Quit(app.ctx)
	})

	if runtime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}

	// --- Format Menu
	FormatMenu := AppMenu.AddSubmenu("Format")
	FormatMenu.AddText("Bold", keys.CmdOrCtrl("B"), func(_ *menu.CallbackData) {
		// TODO: transform selected text to bold
	})
	FormatMenu.AddText("Italic", keys.CmdOrCtrl("I"), func(_ *menu.CallbackData) {
		// TODO: transform selected text to italic
	})
	FormatMenu.AddText("Underline", keys.CmdOrCtrl("U"), func(_ *menu.CallbackData) {
		// TODO: transform selected text to underline
	})
	FormatMenu.AddSeparator()
	FormatMenu.AddText("Strikethrough", keys.CmdOrCtrl("D"), func(_ *menu.CallbackData) {
		// TODO: transform selected text to strikethrough
	})
	FormatMenu.AddSeparator()
	FormatMenu.AddText("Highlight", keys.Combo("H", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *menu.CallbackData) {
		// TODO: transform selected text to highlight
	})
	FormatMenu.AddSeparator()
	FormatMenu.AddText("Inline Code", keys.CmdOrCtrl("`"), func(_ *menu.CallbackData) {
		// TODO: transform selected text to inline code
	})
	FormatMenu.AddText("Inline Math", keys.Combo("M", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *menu.CallbackData) {
		// TODO: transform selected text to inline math
	})
	FormatMenu.AddSeparator()
	FormatMenu.AddText("Hyperlink", keys.CmdOrCtrl("L"), func(_ *menu.CallbackData) {
		// TODO: transform selected text to hyperlink
	})
	FormatMenu.AddText("Image", keys.Combo("I", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *menu.CallbackData) {
		// TODO: open file dialog to insert an image
	})

	return AppMenu
}
