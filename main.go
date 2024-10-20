package main

import (
	"fmt"
	"mousek/infra/base"
	"mousek/infra/config"
	"mousek/infra/keyboardctl"
	"mousek/infra/mousectl"
	"mousek/infra/util"
	"time"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

// 1go:embed frontend/dist
// var assets embed.FS

var vkCodesMulitiSpeedLevelArr = []uint32{keyboardctl.VK_1, keyboardctl.VK_2, keyboardctl.VK_3, keyboardctl.VK_4, keyboardctl.VK_5}

func main() {

	keyboardProcess()
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
// func main2() {

// 	// Create a new Wails application by providing the necessary options.
// 	// Variables 'Name' and 'Description' are for application metadata.
// 	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
// 	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
// 	// 'Mac888898888' options tailor the application when running an macOS.
// 	app := application.New(application.Options{
// 		Name: "mousek",
// 		Windows: application.WindowsOptions{
// 			DisableQuitOnLastWindowClosed: true,
// 		},
// 		Description: "A demo of using raw HTML & CSS",
// 		Services: []application.Service{
// 			application.NewService(&GreetService{}),
// 		},
// 		Assets: application.AssetOptions{
// 			Handler: application.AssetFileServerFS(assets),
// 		},
// 		Mac: application.MacOptions{
// 			ApplicationShouldTerminateAfterLastWindowClosed: true,
// 		},
// 		// Plugins: map[string]application.Plugin{
// 		// 	"start_at_login": start_at_login.NewPlugin(),
// 		// },
// 	})
// 	start_at_login := start_at_login.NewPlugin(start_at_login.Config{
// 		RegistryKey: "mousek.exe",
// 	})
// 	start_at_login.StartAtLogin(true)

// 	InitAppWraper(app)

// 	tray := app.NewSystemTray()
// 	tray.SetLabel("systemtray test")
// 	trayMenu := application.NewMenu()

// 	// TODO remove it
// 	// StartOptionView()

// 	optionMenu := trayMenu.Add("Options")
// 	optionMenu.OnClick(func(ctx *application.Context) {
// 		fmt.Printf("enter option menu \n")
// 		StartOptionView()
// 	})

// 	exitMenuItem := trayMenu.Add("Exit")
// 	exitMenuItem.OnClick(func(ctx *application.Context) {
// 		fmt.Printf("tray menu exit\n")
// 		os.Exit(0)
// 	})

// 	tray.SetMenu(trayMenu)
// 	tray.OnClick(func() {
// 		toggleControlMode()

// 		// fmt.Println("on click system tray")
// 		// fmt.Println(app.CurrentWindow().IsVisible())
// 		// if app.CurrentWindow().IsVisible() {
// 		// 	app.Hide()
// 		// } else {
// 		// 	app.Show()
// 		// }
// 	})

// 	// Create a new window with the necessary options.
// 	// 'Title' is the title of the window.
// 	// 'Mac' options tailor the window when running on macOS.
// 	// 'BackgroundColour' is the background colour of the window.
// 	// 'URL' is the URL that will be loaded into the webview.
// 	// app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
// 	// 	Title: "Options",
// 	// 	Mac: application.MacWindow{
// 	// 		InvisibleTitleBarHeight: 50,
// 	// 		Backdrop:                application.MacBackdropTranslucent,
// 	// 		TitleBar:                application.MacTitleBarHiddenInset,
// 	// 	},
// 	// 	BackgroundColour: application.NewRGB(27, 38, 54),
// 	// 	URL:              "/",
// 	// })

// 	// Create a goroutine that emits an event containing the current time every second.
// 	// The frontend can listen to this event and update the UI accordingly.
// 	go func() {
// 		for {
// 			now := time.Now().Format(time.RFC1123)
// 			app.Events.Emit(&application.WailsEvent{
// 				Name: "time",
// 				Data: now,
// 			})
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	go keyboardProcess()

// 	// Run the application. This blocks until the application has been exited.
// 	err := app.Run()

// 	// If an error occurred while running the application, log it and exit.
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func keyboardProcess() {
	// load config
	config.Init()
	var ts = keyboardctl.GetCodesByNames
	var ta = keyboardctl.GetCodeByName
	settings := config.GetSettings()

	// win+space : activate control mode
	// vkCodesStartControlMode := ts(settings.PresetFunc.ActiveMode)
	// keyboardctl.RegisterNormal(StartControlMode, vkCodesStartControlMode...)

	// alt+r: reset setting
	vkCodesResetSetting := ts(settings.PresetFunc.ResetSetting)
	keyboardctl.RegisterNormal(ResetSetting, vkCodesResetSetting...)

	// alt+0 : toggle control mode
	vkCodesToggleControlMode := ts(settings.PresetFunc.ToggleControlMode)
	keyboardctl.RegisterNormal(ToggleControlMode, vkCodesToggleControlMode...)

	// Q : tmp quit
	vkCodesTmpQuitMode := ts(settings.PresetFunc.TmpQuitMode)
	keyboardctl.RegisterOne(TmpQuitControlMode, vkCodesTmpQuitMode...)

	// space+esc : quit control mode
	// vkCodesQuitControlMode := ts(settings.PresetFunc.QuitMode)
	// keyboardctl.RegisterNormal(QuitControlMode, vkCodesQuitControlMode...)

	// space+comma : open setting panel
	// vkCodesOpenSetting := ts(settings.PresetFunc.OpenSetting)
	// keyboardctl.RegisterOne(ToggleOptionView, vkCodesOpenSetting...)

	// shift + 1\2\3\4\5 : in ModeControl ,control the speed of your mouse scroll
	vkCodesMulitiScrollSpeedLevel := [][]uint32{{keyboardctl.VK_LSHIFT, keyboardctl.VK_1}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_2}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_3}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_4}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_5}}
	keyboardctl.RegisterMulti(ScrollSpeedLevelSwitch, vkCodesMulitiScrollSpeedLevel...)

	// 1\2\3\4\5 : in ModeControl, control the speed of your mouse move
	vkCodesMulitiSpeedLevel := [][]uint32{{keyboardctl.VK_1}, {keyboardctl.VK_2}, {keyboardctl.VK_3}, {keyboardctl.VK_4}, {keyboardctl.VK_5}}
	keyboardctl.RegisterMulti(SpeedLevelSwitch, vkCodesMulitiSpeedLevel...)

	// H\J\K\L : in ModeControl, control the mouse movement like vim
	// W\A\S\D : in ModeControl, control the mouse movement like fps game
	vkCodesSetMousePosUpFast := ts(settings.PresetFunc.MouseMove.Fast.Up)
	vkCodesSetMousePosDownFast := ts(settings.PresetFunc.MouseMove.Fast.Down)
	vkCodesSetMousePosLeftFast := ts(settings.PresetFunc.MouseMove.Fast.Left)
	vkCodesSetMousePosRightFast := ts(settings.PresetFunc.MouseMove.Fast.Right)
	vkCodesSetMousePosUpSlow := ts(settings.PresetFunc.MouseMove.Slow.Up)
	vkCodesSetMousePosDownSlow := ts(settings.PresetFunc.MouseMove.Slow.Down)
	vkCodesSetMousePosLeftSlow := ts(settings.PresetFunc.MouseMove.Slow.Left)
	vkCodesSetMousePosRightSlow := ts(settings.PresetFunc.MouseMove.Slow.Right)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionDown, mousectl.SpeedFast), vkCodesSetMousePosDownFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionUp, mousectl.SpeedFast), vkCodesSetMousePosUpFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionLeft, mousectl.SpeedFast), vkCodesSetMousePosLeftFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionRight, mousectl.SpeedFast), vkCodesSetMousePosRightFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionUp, mousectl.SpeedSlow), vkCodesSetMousePosUpSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionDown, mousectl.SpeedSlow), vkCodesSetMousePosDownSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionLeft, mousectl.SpeedSlow), vkCodesSetMousePosLeftSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionRight, mousectl.SpeedSlow), vkCodesSetMousePosRightSlow...)

	// I\R : in ModeControl, simulate mouse left button click
	// vkCodesMouseLeftClick := [][]uint32{{keyboardctl.VK_I}, {keyboardctl.VK_R}}
	vkCodesMouseLeftClick := [][]uint32{{ta(settings.PresetFunc.MouseLeftButtonClick.Primary)}, {ta(settings.PresetFunc.MouseLeftButtonClick.Secondary)}}
	keyboardctl.RegisterMulti(MouseLeftClick, vkCodesMouseLeftClick...)

	// O\T : in ModeControl, simulate mouse right button click
	// vkCodesMouseRightClick := [][]uint32{{keyboardctl.VK_O}, {keyboardctl.VK_T}}
	vkCodesMouseRightClick := [][]uint32{{ta(settings.PresetFunc.MouseRightButtonClick.Primary)}, {ta(settings.PresetFunc.MouseRightButtonClick.Secondary)}}
	keyboardctl.RegisterMulti(MouseRightClick, vkCodesMouseRightClick...)

	// C\N : in ModeControl, simulate mouse left button hold
	// vkCoodesLeftDown := [][]uint32{{keyboardctl.VK_C}, {keyboardctl.VK_N}}
	vkCoodesLeftDown := [][]uint32{{ta(settings.PresetFunc.MouseLeftButtonHold.Primary)}, {ta(settings.PresetFunc.MouseLeftButtonHold.Secondary)}}
	keyboardctl.RegisterWithReleaseEventMulti(MouseLeftDown, vkCoodesLeftDown...)

	// shift + H\J\K\L : in ModeControl, control the mouse scroll like vim
	// shift + W\A\S\D : in ModeControl, control the mouse scroll like fps game
	vkCodesMouseVerticalScrollDownFast := ts(settings.PresetFunc.MouseScroll.Fast.Down)
	vkCodesMouseVerticalScrollUpFast := ts(settings.PresetFunc.MouseScroll.Fast.Up)
	vkCodesMouseHorizontalScrollLeftFast := ts(settings.PresetFunc.MouseScroll.Fast.Left)
	vkCodesMouseHorizontalScrollRightFast := ts(settings.PresetFunc.MouseScroll.Fast.Right)
	vkCodesMouseVerticalScrollDownSlow := ts(settings.PresetFunc.MouseScroll.Slow.Down)
	vkCodesMouseVerticalScrollUpSlow := ts(settings.PresetFunc.MouseScroll.Slow.Up)
	vkCodesMouseHorizontalScrollLeftSlow := ts(settings.PresetFunc.MouseScroll.Slow.Left)
	vkCodesMouseHorizontalScrollRightSlow := ts(settings.PresetFunc.MouseScroll.Slow.Right)
	// vkCodesMouseVerticalScrollUpFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_K}
	// vkCodesMouseHorizontalScrollLeftFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_H}
	// vkCodesMouseHorizontalScrollRightFast := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_L}
	// vkCodesMouseVerticalScrollDownSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_S}
	// vkCodesMouseVerticalScrollUpSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_W}
	// vkCodesMouseHorizontalScrollLeftSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_A}
	// vkCodesMouseHorizontalScrollRightSlow := []uint32{keyboardctl.VK_LSHIFT, keyboardctl.VK_D}
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

func toggleControlMode() {
	base.SetMode(1 - base.GetMode())
	fmt.Printf("toggle mode to:%d\n", base.GetMode())
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
}

// func ToggleOptionView(wParam uintptr, vkCode, scanCode uint32) uintptr {
// 	if base.ToggleOptionViewState() {
// 		StartOptionView()
// 	} else {
// 		HideOptionView()
// 	}
// 	return 1
// }

func StartControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d\n", base.GetMode())
	if base.GetMode() == base.ModeControl {
		fmt.Println("already in control mode", time.Now())
	} else {
		base.SetMode(base.ModeControl)
		fmt.Println("change to control mode", time.Now())
	}
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	return 1
}

func ResetSetting(wParam uintptr, vkCode, scanCode uint32) uintptr {

	fmt.Printf("user restore setting\n")
	config.RestoreSettings()
	return 1
}

func QuitControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d\n", base.GetMode())
	fmt.Println()
	if base.GetMode() == base.ModeControl {
		fmt.Println("change to normal mode", time.Now())
		base.SetMode(base.ModeNormal)
	} else {
		fmt.Println("already in normal mode", time.Now())
	}
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	return 0
}

func ScrollSpeedLevelSwitch(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// if base.GetMode() != base.ModeControl {
	// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	if util.Contains[uint32](vkCodesMulitiSpeedLevelArr, uint32(vkCode)) {
		speedLevel := int(vkCode) - keyboardctl.VK_1 + 1
		base.SetScrollSpeedLevel(speedLevel)
		fmt.Printf("change speed to :%d\n", base.GetScrollSpeed())
	}
	return 1
}

func SpeedLevelSwitch(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// if base.GetMode() != base.ModeControl {
	// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	if util.Contains[uint32](vkCodesMulitiSpeedLevelArr, uint32(vkCode)) {
		speedLevel := int(vkCode) - keyboardctl.VK_1 + 1
		base.SetMoveSpeedLevel(speedLevel)
		fmt.Printf("change speed to :%d\n", base.GetMoveSpeedLevel())
	}
	return 1
}

func MouseLeftClick(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d\n", base.GetMode())
	// if base.GetMode() != base.ModeControl {
	// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	mousectl.LeftClick()
	fmt.Printf("mouse left click\n")
	return 1
}

func MouseRightClick(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d\n", base.GetMode())
	// if base.GetMode() != base.ModeControl {
	// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	mousectl.RightClick()
	fmt.Printf("mouse right click\n")
	return 1
}

func MouseLeftDown(wParam uintptr, vkCode, scanCode uint32) uintptr {
	if wParam == keyboardctl.WM_KEYDOWN {
		fmt.Printf("mouse left button down\n")
		mousectl.MouseLeftDown()
	} else if wParam == keyboardctl.WM_KEYUP {
		fmt.Printf("mouse left button up\n")
		mousectl.MouseLeftUp()
	}
	return 1
}

func MoveMouseFunc(direction mousectl.MoveDirection, speedType mousectl.MoveSpeedType) keyboardctl.Callback2 {
	return func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		// if base.GetMode() != base.ModeControl {
		// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
		// 	return 0
		// }
		mousectl.MoveMouseCtrl(direction, speedType)
		return 1
	}
}

func ScrollMouseFunc(direction mousectl.ScrollDirection, speed mousectl.MoveSpeedType) keyboardctl.Callback2 {
	return func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		// if base.GetMode() != base.ModeControl {
		// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
		// 	return 0
		// }
		mousectl.ScrollMouseCtrl(direction, speed)
		return 1
	}
}

func TmpQuitControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d\n", base.GetMode())
	fmt.Println()
	if base.GetMode() == base.ModeControl {
		fmt.Println("change to normal mode", time.Now())
		base.SetMode(base.ModeNormal)
	} else {
		fmt.Println("already in normal mode", time.Now())
	}
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	return 0
}

func ToggleControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	toggleControlMode()
	return 1
}
