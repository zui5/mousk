package ui

import (
	"fmt"
	"mousk/common/logger"
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
		// pureTextDialogWindow.SetHTML(fmt.Sprintf("<div style=\"background-color:white\">%s</div>", text))
		html := fmt.Sprintf(`
			<html>
			<head>
				<style>
					body {
						margin: 0;
						padding: 0;
						background-color: rgba(0, 0, 0, 0.7);
						color: white;
						font-size: 20px;
						display: flex;
						align-items: center;
						justify-content: center;
						height: 100%%;
					}
				</style>
			</head>
			<body>
				%s
			</body>
			</html>`, text)

		pureTextDialogWindow.SetHTML(html)
		pureTextDialogWindow.Show()
		time.Sleep(500 * time.Millisecond)
		pureTextDialogWindow.Hide()
	}()
}

func initPureTextDialogWindow() *application.WebviewWindow {

	diaglogView := application.NewWindow(application.WebviewWindowOptions{
		Name:           "dialog",
		Width:          350,
		Height:         50,
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
	logger.Infof("", "fuck notificationview：%+v", diaglogView)
	return diaglogView
}
