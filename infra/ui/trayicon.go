package ui

import (
	"mousk/common/logger"
	"mousk/infra/base"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func initTray() {
	TrayInstance.SetLabel("system test")
	TrayInstance.SetDarkModeIcon(GetTrayIcon(base.GetMode()))
	TrayInstance.SetIcon(GetTrayIcon(base.GetMode()))
}

func ToggleTrayIcon() {
	TrayInstance.SetDarkModeIcon(GetTrayIcon(base.GetMode()))
	TrayInstance.SetIcon(GetTrayIcon(base.GetMode()))
}

func TrayMenu() {

	trayMenu := application.NewMenu()
	optionMenu := trayMenu.Add("Options")
	optionMenu.OnClick(func(ctx *application.Context) {
		logger.Infof("", "enter option menu ")
		// StartOptionView()
	})

	exitMenuItem := trayMenu.Add("Exit")
	exitMenuItem.OnClick(func(ctx *application.Context) {
		logger.Infof("", "tray menu exit")
		os.Exit(0)
	})

	TrayInstance.SetMenu(trayMenu)
	TrayInstance.OnClick(func() {
		// toggleControlMode()
		TrayInstance.SetIcon(GetTrayIcon(base.GetMode()))
		TrayInstance.SetDarkModeIcon(GetTrayIcon(base.GetMode()))
	})

}
