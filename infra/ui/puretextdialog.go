package ui

import (
	"fmt"
	"mousk/common/logger"
	"mousk/infra/base"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var pureTextDialogWindow *application.WebviewWindow = nil
var helperDialogWindow *application.WebviewWindow = nil

func Message(text string) {
	if pureTextDialogWindow == nil {
		pureTextDialogWindow = initPureTextDialogWindow()
		// pureTextDialogWindow.Hide()
	}
	go func() {
		// pureTextDialogWindow.SetHTML(fmt.Sprintf("<div style="background-color:white\">%s</div>", text))
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

func initHelperDialogWindow() *application.WebviewWindow {
	diaglogView := application.NewWindow(application.WebviewWindowOptions{
		Name:              "helper",
		Width:             1000,
		Height:            750,
		AlwaysOnTop:       true,
		EnableDragAndDrop: true,
		DisableResize:     true,
		Frameless:         true,
		Centered:          true,
		BackgroundType:    application.BackgroundTypeSolid,
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

func ToggleHelper(text string, helpmode int) {
	if helperDialogWindow == nil {
		helperDialogWindow = initHelperDialogWindow()
		// pureTextDialogWindow.Hide()
	}
	go func() {
		// pureTextDialogWindow.SetHTML(fmt.Sprintf("<div style="background-color:white\">%s</div>", text))
		html := `<html>
<head>
  <style>
    body {
      margin: 0;
      padding: 0;
      background-color: #121212;
      color: white;
      font-family: 'Arial', sans-serif;
      font-size: 16px;
      display: flex;
      align-items: center;
      justify-content: center;
      height: 100vh;
      line-height: 1.5;
    }

    .container {
      background-color: #1c1c1c;
      padding: 30px;
      border-radius: 10px;
      width: 100%;
      max-width: 900px;
      box-shadow: 0 4px 20px rgba(0, 0, 0, 0.7);
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
    }

    .section {
      flex: 1;
      min-width: 300px;
      margin-right: 20px;
    }

    h2 {
      font-size: 22px;
      margin-bottom: 15px;
      color: #f0f0f0;
      border-bottom: 2px solid #444;
      padding-bottom: 5px;
    }

    ul {
      padding-left: 20px;
      margin-bottom: 20px;
    }

    li {
      margin-bottom: 10px;
    }

    code {
      background-color: #333;
      color: #ffcc00;
      padding: 3px 8px;
      border-radius: 5px;
      font-size: 1.1em;
      font-weight: bold;
    }

    /* Quick key emphasis */
    .shortcut {
      color: #00b0ff;
      font-weight: bold;
      background-color: #333;
      padding: 4px 8px;
      border-radius: 5px;
    }

    /* Remove the margin for the last item in each section */
    .section ul li:last-child {
      margin-bottom: 0;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="section">
      <h2>General Controls:</h2>
      <ul>
        <li>Toggle Mode: <code class="shortcut">LALT + 0</code></li>
        <li>Open Settings: <code class="shortcut">SPACE + COMMA</code></li>
        <li>Reset Settings: <code class="shortcut">LALT + R</code></li>
        <li>Force Quit: <code class="shortcut">LCONTROL + LSHIFT + A</code></li>
        <li>Temporary Quit Control Mode (Only in control mode): <code class="shortcut">Q</code></li>
      </ul>

      <h2>Mouse Movement:</h2>
      <ul>
        <li>Fast mode: Use <code class="shortcut">J</code>, <code class="shortcut">H</code>, <code class="shortcut">L</code>, <code class="shortcut">K</code> for down, left, right, and up respectively.</li>
        <li>Slow mode: Use <code class="shortcut">S</code>, <code class="shortcut">A</code>, <code class="shortcut">D</code>, <code class="shortcut">W</code> for down, left, right, and up respectively.</li>
        <li><strong>Speed levels:</strong> Use <code class="shortcut">1</code>, <code class="shortcut">2</code>, <code class="shortcut">3</code>, <code class="shortcut">4</code>, <code class="shortcut">5</code> to switch between speed levels, <code class="shortcut">5</code> is highest.</li>
      </ul>
    </div>

    <div class="section">
      <h2>Mouse Scroll Simulation:</h2>
      <ul>
        <li>Fast mode: Use <code class="shortcut">LSHIFT + J</code>, <code class="shortcut">LSHIFT + H</code>, <code class="shortcut">LSHIFT + L</code>, <code class="shortcut">LSHIFT + K</code> for down, left, right, and up respectively.</li>
        <li>Slow mode: Use <code class="shortcut">LSHIFT + S</code>, <code class="shortcut">LSHIFT + A</code>, <code class="shortcut">LSHIFT + D</code>, <code class="shortcut">LSHIFT + W</code> for down, left, right, and up respectively.</li>
        <li><strong>Speed levels:</strong> Use <code class="shortcut">Shift + 1</code>, <code class="shortcut">Shift + 2</code>, <code class="shortcut">Shift + 3</code>, <code class="shortcut">Shift + 4</code>, <code class="shortcut">Shift + 5</code> to switch between speed levels, <code class="shortcut">Shift + 5</code> is highest.</li>
      </ul>

      <h2>Mouse Clicks:</h2>
      <ul>
        <li>Left Button Click: Primary - <code class="shortcut">I</code>, Secondary - <code class="shortcut">R</code>.</li>
        <li>Right Button Click: Primary - <code class="shortcut">O</code>, Secondary - <code class="shortcut">T</code>.</li>
        <li>Left Button Hold (C/N hold mouse, then H/J/K/L do the movement): Primary - <code class="shortcut">C</code>, Secondary - <code class="shortcut">N</code>.</li>
      </ul>
    </div>
  </div>
</body>
</html>
`
		if helpmode == 1 {
			base.SetHelperMode(1)
			helperDialogWindow.SetHTML(html)
			helperDialogWindow.Show()
			time.Sleep(5000 * time.Millisecond)
			helperDialogWindow.Hide()
			base.SetHelperMode(0)
		} else {
			helperDialogWindow.Hide()
			base.SetHelperMode(0)
		}

	}()
}
