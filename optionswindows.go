package main

import (
	"mousk/common/logger"
	"mousk/infra/ui"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// var
var optionsView *application.WebviewWindow

func StartOptionView() {
	if ui.AppInstance == nil {
		logger.Infof("", "appwraper not initialized")
		return
	}
	if optionsView == nil {
		optionsView = ui.AppInstance.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
			Name:              "Options",
			Title:             "Options",
			Width:             800,
			Height:            600,
			AlwaysOnTop:       false,
			URL:               "#/option",
			DisableResize:     true,
			EnableDragAndDrop: true,
			Frameless:         false,
			MinWidth:          800,
			MinHeight:         400,
			MaxWidth:          1200,
			MaxHeight:         800,
			StartState:        0,
			Centered:          true,
			// BackgroundType:   application.BackgroundTypeTranslucent,
			BackgroundType:   application.BackgroundTypeSolid,
			BackgroundColour: application.NewRGB(0, 0, 0),
			// BackgroundColour:        application.RGBA{},
			// HTML:                    "",
			// JS:                      "",
			// CSS:                     "",
			// X:                       0,
			// Y:                       0,
			// FullscreenButtonEnabled: false,
			// Hidden:                  false,
			// Zoom:                    0,
			ZoomControlEnabled: true,
			// OpenInspectorOnStartup:  false,
			// Mac:                     application.MacWindow{},
			// Windows:                 application.WindowsWindow{},
			// Linux:                   application.LinuxWindow{},
			ShouldClose: func(window *application.WebviewWindow) bool {
				window.Hide()
				logger.Infof("", "view show close")
				return false
			},
			// DevToolsEnabled:            false,
			// DefaultContextMenuDisabled: false,
			// KeyBindings:                map[string]func(window *application.WebviewWindow){},
			// IgnoreMouseEvents:          false,
		})
		logger.Infof("", "fuck optionview：%+v", optionsView)
	}
	optionsView.Show()
}

func HideOptionView() {
	if ui.AppInstance == nil {
		logger.Infof("", "appwraper not initialized")
		return
	}
	if optionsView == nil {
		return
	}
	optionsView.Hide()
}
