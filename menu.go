package main

import (
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

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

	// --- Transform Menu
	TransformMenu := AppMenu.AddSubmenu("Transform")
	TransformMenu.AddText("Heading 1", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		// TODO: transform selected text
	})
	TransformMenu.AddText("Heading 2", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		// TODO: transform selected text
	})
	TransformMenu.AddText("Heading 3", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		// TODO: transform selected text
	})
	TransformMenu.AddText("Heading 4", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		// TODO: transform selected text
	})
	TransformMenu.AddText("Heading 5", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		// TODO: transform selected text
	})
	TransformMenu.AddText("Heading 6", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		// TODO: transform selected text
	})

	// --- Window Menu
	WindowMenu := AppMenu.AddSubmenu("Window")
	WindowMenu.AddText("Minimize", keys.CmdOrCtrl("M"), func(_ *menu.CallbackData) {
		wailsRuntime.WindowMinimise(app.ctx)
	})
	WindowMenu.AddText("Maximize", nil, func(_ *menu.CallbackData) {
		wailsRuntime.WindowMaximise(app.ctx)
	})
	WindowMenu.AddCheckbox("Always On Top", false, nil, func(cb *menu.CallbackData) {
		wailsRuntime.WindowSetAlwaysOnTop(app.ctx, cb.MenuItem.Checked)
	})
	WindowMenu.AddSeparator()
	WindowMenu.AddText("Fullscreen", nil, func(_ *menu.CallbackData) {
		wailsRuntime.WindowFullscreen(app.ctx)
	})
	WindowMenu.AddText("Unfulscreen", nil, func(_ *menu.CallbackData) {
		wailsRuntime.WindowUnfullscreen(app.ctx)
	})

	// --- Themes Menu
	ThemesMenu := AppMenu.AddSubmenu("Themes")
	ThemesMenu.AddText("Axolotl - Light", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "change-theme", "axolotl-light")
	})
	ThemesMenu.AddText("Scarecrow - Dark", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "change-theme", "scarecrow-dark")
	})

	// --- Help Menu
	HelpMenu := AppMenu.AddSubmenu("Help")
	HelpMenu.AddText("Quick Start", nil, func(_ *menu.CallbackData) {
		wailsRuntime.BrowserOpenURL(app.ctx, "https://github.com/dmytro-feshchenko/memo-craft?tab=readme-ov-file#how-to-build")
	})
	HelpMenu.AddText("Watch on GitHub", nil, func(_ *menu.CallbackData) {
		wailsRuntime.BrowserOpenURL(app.ctx, "https://github.com/dmytro-feshchenko/memo-craft")
	})

	return AppMenu
}
