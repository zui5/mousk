package ui

import "github.com/wailsapp/wails/v3/pkg/application"

type AppWrapper struct{ *application.App }
type TrayWrapper struct{ *application.SystemTray }

// var optionView string
var AppInstance *AppWrapper
var TrayInstance *TrayWrapper

func InitWrapper(app *application.App) {
	AppInstance = &AppWrapper{
		App: app,
	}
	TrayInstance = &TrayWrapper{app.NewSystemTray()}
	initTray()
}
