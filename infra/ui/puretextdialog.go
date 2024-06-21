package ui

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var pureTextDialogWindow *application.WebviewWindow = nil

func Message(text string) {
	if pureTextDialogWindow == nil {
		pureTextDialogWindow = initPureTextDialogWindow()
		// pureTextDialogWindow.Hide()
	}
	go func() {
		pureTextDialogWindow.SetHTML(fmt.Sprintf("<div style=\"background-color:white\">%s</div>", text))
		pureTextDialogWindow.Show()
		time.Sleep(2 * time.Second)
		pureTextDialogWindow.Hide()
	}()
}

func initPureTextDialogWindow() *application.WebviewWindow {
	return application.NewWindow(application.WebviewWindowOptions{
		Name:                    "111",
		Width:                   200,
		Height:                  50,
		AlwaysOnTop:             true,
		DisableResize:           true,
		Frameless:               true,
		HTML:                    "<div style=\"background-color:yellow\">fucking text</div>",
		Centered:                true,
		BackgroundType:          application.BackgroundTypeTransparent,
		BackgroundColour:        application.RGBA{},
		FullscreenButtonEnabled: false,
		Windows: application.WindowsWindow{
			DisableMinimiseButton:             true,
			DisableMaximiseButton:             true,
			WebviewGpuIsDisabled:              true,
			DisableIcon:                       true,
			DisableMenu:                       true,
			DisableFramelessWindowDecorations: true,
			HiddenOnTaskbar:                   true,
		},
		// DevToolsEnabled: true,
		// OpenInspectorOnStartup:     true,
		Linux:                      application.LinuxWindow{},
		DefaultContextMenuDisabled: true,
	})
}
