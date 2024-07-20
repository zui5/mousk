package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"mousek/infra/keyboardctl"
	"mousek/infra/mousectl"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/plugins/experimental/start_at_login"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

var vkCodesMulitiSpeedLevelArr = []uint32{keyboardctl.VK_1, keyboardctl.VK_2, keyboardctl.VK_3, keyboardctl.VK_4, keyboardctl.VK_5}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac888898888' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name: "mousek",
		Windows: application.WindowsOptions{
			DisableQuitOnLastWindowClosed: true,
		},
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		// Plugins: map[string]application.Plugin{
		// 	"start_at_login": start_at_login.NewPlugin(),
		// },
	})
	start_at_login := start_at_login.NewPlugin(start_at_login.Config{
		RegistryKey: "mousek.exe",
	})
	start_at_login.StartAtLogin(true)

	InitAppWraper(app)

	tray := app.NewSystemTray()
	tray.SetLabel("systemtray test")
	trayMenu := application.NewMenu()

	// TODO remove it
	// StartOptionView()

	optionMenu := trayMenu.Add("Options")
	optionMenu.OnClick(func(ctx *application.Context) {
		fmt.Printf("enter option menu \n")
		StartOptionView()
	})

	exitMenuItem := trayMenu.Add("Exit")
	exitMenuItem.OnClick(func(ctx *application.Context) {
		fmt.Printf("tray menu exit\n")
		os.Exit(0)
	})

	tray.SetMenu(trayMenu)
	tray.OnClick(func() {
		ToggleControlMode()

		// fmt.Println("on click system tray")
		// fmt.Println(app.CurrentWindow().IsVisible())
		// if app.CurrentWindow().IsVisible() {
		// 	app.Hide()
		// } else {
		// 	app.Show()
		// }
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	// app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
	// 	Title: "Options",
	// 	Mac: application.MacWindow{
	// 		InvisibleTitleBarHeight: 50,
	// 		Backdrop:                application.MacBackdropTranslucent,
	// 		TitleBar:                application.MacTitleBarHiddenInset,
	// 	},
	// 	BackgroundColour: application.NewRGB(27, 38, 54),
	// 	URL:              "/",
	// })

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Events.Emit(&application.WailsEvent{
				Name: "time",
				Data: now,
			})
			time.Sleep(time.Second)
		}
	}()

	go keyboardProcess()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}

func keyboardProcess() {

	// win+space : activate control mode
	// vkCodesWinSpace := []uint32{keyboardctl.VK_LWIN, keyboardctl.VK_SPACE}
	vkCodesTabSpace := []uint32{keyboardctl.VK_TAB, keyboardctl.VK_SPACE}
	keyboardctl.RegisterOne(StartControlMode, vkCodesTabSpace...)
	// keyboardctl.RegisterOne(StartControlMode, vkCodesWinSpace...)

	vkCodesSpaceComma := []uint32{keyboardctl.VK_SPACE, keyboardctl.VK_OEM_COMMA}
	keyboardctl.RegisterOne(ToggleOptionView, vkCodesSpaceComma...)

	// space+esc : quit control mode
	vkCodesEsc := []uint32{keyboardctl.VK_ESCAPE, keyboardctl.VK_ESCAPE}
	keyboardctl.RegisterOne(QuitControlMode, vkCodesEsc...)

	// shift + 1\2\3\4\5 : in ModeControl ,control the speed of your mouse scroll
	vkCodesMulitiScrollSpeedLevel := [][]uint32{{keyboardctl.VK_LSHIFT, keyboardctl.VK_1}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_2}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_3}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_4}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_5}}
	keyboardctl.RegisterMulti(ScrollSpeedLevelSwitch, vkCodesMulitiScrollSpeedLevel...)

	// 1\2\3\4\5 : in ModeControl, control the speed of your mouse move
	vkCodesMulitiSpeedLevel := [][]uint32{{keyboardctl.VK_1}, {keyboardctl.VK_2}, {keyboardctl.VK_3}, {keyboardctl.VK_4}, {keyboardctl.VK_5}}
	keyboardctl.RegisterMulti(SpeedLevelSwitch, vkCodesMulitiSpeedLevel...)

	// H\J\K\L : in ModeControl, control the mouse movement like vim
	// W\A\S\D : in ModeControl, control the mouse movement like fps game
	vkCodesSetMousePosUpFast := []uint32{keyboardctl.VK_K}
	vkCodesSetMousePosDownFast := []uint32{keyboardctl.VK_J}
	vkCodesSetMousePosLeftFast := []uint32{keyboardctl.VK_H}
	vkCodesSetMousePosRightFast := []uint32{keyboardctl.VK_L}
	vkCodesSetMousePosUpSlow := []uint32{keyboardctl.VK_W}
	vkCodesSetMousePosDownSlow := []uint32{keyboardctl.VK_S}
	vkCodesSetMousePosLeftSlow := []uint32{keyboardctl.VK_A}
	vkCodesSetMousePosRightSlow := []uint32{keyboardctl.VK_D}
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionDown, mousectl.SpeedFast), vkCodesSetMousePosDownFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionUp, mousectl.SpeedFast), vkCodesSetMousePosUpFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionLeft, mousectl.SpeedFast), vkCodesSetMousePosLeftFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionRight, mousectl.SpeedFast), vkCodesSetMousePosRightFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionUp, mousectl.SpeedSlow), vkCodesSetMousePosUpSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionDown, mousectl.SpeedSlow), vkCodesSetMousePosDownSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionLeft, mousectl.SpeedSlow), vkCodesSetMousePosLeftSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionRight, mousectl.SpeedSlow), vkCodesSetMousePosRightSlow...)

	// I\R : in ModeControl, simulate mouse left button click
	vkCodesMouseLeftClick := [][]uint32{{keyboardctl.VK_I}, {keyboardctl.VK_R}}
	keyboardctl.RegisterMulti(MouseLeftClick, vkCodesMouseLeftClick...)

	// O\T : in ModeControl, simulate mouse right button click
	vkCodesMouseRightClick := [][]uint32{{keyboardctl.VK_O}, {keyboardctl.VK_T}}
	keyboardctl.RegisterMulti(MouseRightClick, vkCodesMouseRightClick...)

	// C\N : in ModeControl, simulate mouse left button hold
	vkCoodesLeftDown := [][]uint32{{keyboardctl.VK_C}, {keyboardctl.VK_N}}
	keyboardctl.RegisterWithReleaseEventMulti(MouseLeftDown, vkCoodesLeftDown...)

	// shift + H\J\K\L : in ModeControl, control the mouse scroll like vim
	// shift + W\A\S\D : in ModeControl, control the mouse scroll like fps game
	vkCodesMouseVerticalScrollDownFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_J}
	vkCodesMouseVerticalScrollUpFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_K}
	vkCodesMouseHorizontalScrollLeftFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_H}
	vkCodesMouseHorizontalScrollRightFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_L}
	vkCodesMouseVerticalScrollDownSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_S}
	vkCodesMouseVerticalScrollUpSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_W}
	vkCodesMouseHorizontalScrollLeftSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_A}
	vkCodesMouseHorizontalScrollRightSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_D}
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalDown, mousectl.SpeedFast), vkCodesMouseVerticalScrollDownFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalUp, mousectl.SpeedFast), vkCodesMouseVerticalScrollUpFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalLeft, mousectl.SpeedFast), vkCodesMouseHorizontalScrollLeftFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalRight, mousectl.SpeedFast), vkCodesMouseHorizontalScrollRightFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalLeft, mousectl.SpeedSlow), vkCodesMouseHorizontalScrollLeftSlow...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalRight, mousectl.SpeedSlow), vkCodesMouseHorizontalScrollRightSlow...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalDown, mousectl.SpeedSlow), vkCodesMouseVerticalScrollDownSlow...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalUp, mousectl.SpeedSlow), vkCodesMouseVerticalScrollUpSlow...)

	// _____________________________________________________________________________________________________________________________11111111111111111111111111222222222222222222222222222222222223333333333333333
	// main keyboard event listener
	keyboardctl.RawKeyboardListener(keyboardctl.LowLevelKeyboardCallback)

}

// // 控制鼠标在指定显示器的四周移动
// func moveMouseAround(monitor monitor.RECT) {
// 	x := int(monitor.Left)
// 	y := int(monitor.Top)

// 	/* 	width := int(monitor.Right - monitor.Left)
// 	   	height := int(monitor.Bottom - monitor.Top)
// 	*/
// 	// 向右移动到显示器右边缘
// 	for x < int(monitor.Right) {
// 		mousectl.SetMousePos(x, y)
// 		x += 10
// 		time.Sleep(5 * time.Millisecond)
// 	}

// 	// 向下移动到显示器底边缘
// 	for y < int(monitor.Bottom) {
// 		mousectl.SetMousePos(x, y)
// 		y += 10
// 		time.Sleep(5 * time.Millisecond)
// 	}

// 	// 向左移动到显示器左边缘
// 	for x > int(monitor.Left) {
// 		mousectl.SetMousePos(x, y)
// 		x -= 10
// 		time.Sleep(5 * time.Millisecond)
// 	}

// 	// 向上移动到显示器上边缘
// 	for y > int(monitor.Top) {
// 		mousectl.SetMousePos(x, y)
// 		y -= 10
// 		time.Sleep(5 * time.Millisecond)
// 	}
// }
