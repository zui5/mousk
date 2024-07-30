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
	diaglogView := application.NewWindow(application.WebviewWindowOptions{
		Name:           "dialog",
		Width:          300,
		Height:         40,
		AlwaysOnTop:    true,
		DisableResize:  true,
		Frameless:      true,
		Centered:       true,
		BackgroundType: application.BackgroundTypeSolid,
		BackgroundColour: application.RGBA{
			Red:   255,
			Green: 255,
			Blue:  255,
			Alpha: 1,
		},
		FullscreenButtonEnabled: false,
	})
	fmt.Printf("fuck notificationview：%+v\n", diaglogView)
	return diaglogView
}
