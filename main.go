package main

import (
	"embed"
	"fmt"
	"log"
	"mousk/common/logger"
	"mousk/infra/base"
	"mousk/infra/config"
	"mousk/infra/keyboardctl"
	"mousk/infra/mousectl"
	"mousk/infra/ui"
	"mousk/infra/util"
	"mousk/service"
	"os"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

var vkCodesMulitiSpeedLevelArr = []uint32{keyboardctl.VK_1, keyboardctl.VK_2, keyboardctl.VK_3, keyboardctl.VK_4, keyboardctl.VK_5}

func init() {

}

func main() {
	app := application.New(application.Options{
		Name: "mousk",
		Windows: application.WindowsOptions{
			DisableQuitOnLastWindowClosed: true,
		},
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{}),
			application.NewService(&service.StartupService{}),
			application.NewService(&service.KeymapService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	ui.InitWrapper(app)

	trayMenu := application.NewMenu()
	optionMenu := trayMenu.Add("Options")
	optionMenu.OnClick(func(ctx *application.Context) {
		logger.Infof("", "enter option menu ")
		StartOptionView()
	})

	exitMenuItem := trayMenu.Add("Exit")
	exitMenuItem.OnClick(func(ctx *application.Context) {
		logger.Infof("", "tray menu exit")
		os.Exit(0)
	})
	tray := ui.TrayInstance
	tray.SetMenu(trayMenu)
	tray.OnClick(func() {
		toggleControlMode()
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tsilor the window when running on macOS.
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

	// Create a goroutine that emits an event contsining the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.EmitEvent("time", now)
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

// main function serves as the application's entry point. It initializes the application, creates a window,
// and stsrts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
// func main2() {

// 	// Create a new Wails application by providing the necessary options.
// 	// Variables 'Name' and 'Description' are for application metsdats.
// 	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
// 	// 'Bind' is a list of Go struct instsnces. The frontend has access to the methods of these instsnces.
// 	// 'Mac888898888' options tsilor the application when running an macOS.
// 	app := application.New(application.Options{
// 		Name: "mousk",
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
// 		// 	"stsrt_at_login": stsrt_at_login.NewPlugin(),
// 		// },
// 	})
// 	stsrt_at_login := stsrt_at_login.NewPlugin(stsrt_at_login.Config{
// 		RegistryKey: "mousk.exe",
// 	})
// 	stsrt_at_login.StartstLogin(true)

// 	InitsppWraper(app)

// 	tray := app.NewSystemTray()
// 	tray.SetLabel("systemtray test")
// 	trayMenu := application.NewMenu()

// 	// TODO remove it
// 	// StartOptionView()

// 	optionMenu := trayMenu.Add("Options")
// 	optionMenu.OnClick(func(ctx *application.Context) {
// 		logger.Infof("","enter option menu ")
// 		StartOptionView()
// 	})

// 	exitMenuItem := trayMenu.Add("Exit")
// 	exitMenuItem.OnClick(func(ctx *application.Context) {
// 		logger.Infof("","tray menu exit")
// 		os.Exit(0)
// 	})

// 	tray.SetMenu(trayMenu)
// 	tray.OnClick(func() {
// 		toggleControlMode()

// 		// logger.Infof("","on click system tray")
// 		// logger.Infof("",app.CurrentWindow().IsVisible())
// 		// if app.CurrentWindow().IsVisible() {
// 		// 	app.Hide()
// 		// } else {
// 		// 	app.Show()
// 		// }
// 	})

// 	// Create a new window with the necessary options.
// 	// 'Title' is the title of the window.
// 	// 'Mac' options tsilor the window when running on macOS.
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

// 	// Create a goroutine that emits an event contsining the current time every second.
// 	// The frontend can listen to this event and update the UI accordingly.
// 	go func() {
// 		for {
// 			now := time.Now().Format(time.RFC1123)
// 			app.Events.Emit(&application.WailsEvent{
// 				Name: "time",
// 				Dats: now,
// 			})
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	go keyboardProcess()

// 	// Run the application. This blocks until the application has been exited.
// 	err := app.Run()

// 	// If an error occurred while running the application, log it and exit.
// 	if err != nil {
// 		log.Fatsl(err)
// 	}
// }

func keyboardProcess() {
	// load config
	// config.Init()
	var ts = keyboardctl.GetCodesByNames
	// var ts = keyboardctl.GetCodeByName
	settings := config.GetSettings()

	// alt+0 : toggle control mode
	vkCodesToggleControlMode := ts(settings.Shortcuts.ActiveMode.Keys)
	keyboardctl.RegisterNormal(ToggleControlMode, 0, vkCodesToggleControlMode...)

	// alt+periot : help pane
	vkCodesHelpPane := ts(settings.Shortcuts.HelpPane.Keys)
	keyboardctl.RegisterNormal(ToggleHelpPane, 0, vkCodesHelpPane...)
	// keyboardctl.RegisterWithReleaseEventMulti(cb keyboardctl.Callback2, prority int, mulitiVkCodes ...[]uint32)

	// 屏蔽部分按键
	keyboardctl.RegisterMulti(BlockKey, -1, keyboardctl.ExportAllCodes()...)

	vkCodesForceQuit := ts(settings.Shortcuts.ForceQuit.Keys)
	keyboardctl.RegisterNormal(ForceQuit, 0, vkCodesForceQuit...)

	// alt+r: reset setting
	vkCodesResetSetting := ts(settings.Shortcuts.ResetSetting.Keys)
	keyboardctl.RegisterNormal(ResetSetting, 0, vkCodesResetSetting...)

	// Q : tmp quit
	vkCodesTmpQuitMode := ts(settings.Shortcuts.TmpQuitMode.Keys)
	keyboardctl.RegisterOne(TmpQuitControlMode, 0, vkCodesTmpQuitMode...)

	// alt+comma : open setting panel
	vkCodesOpenSetting := ts(settings.Shortcuts.OpenSetting.Keys)
	keyboardctl.RegisterNormal(ToggleOptionView, 0, vkCodesOpenSetting...)

	// 1\2\3\4\5 : in ModeControl, control the speed of your mouse move
	vkCodesMulitiSpeedLevel := [][]uint32{{keyboardctl.VK_1}, {keyboardctl.VK_2}, {keyboardctl.VK_3}, {keyboardctl.VK_4}, {keyboardctl.VK_5}}
	keyboardctl.RegisterMulti(SpeedLevelSwitch, 0, vkCodesMulitiSpeedLevel...)

	// H\J\K\L : in ModeControl, control the mouse movement like vim
	// W\A\S\D : in ModeControl, control the mouse movement like fps game
	vkCodesSetMousePosUpFast := ts(settings.Shortcuts.MouseMoveFastUp.Keys)
	vkCodesSetMousePosDownFast := ts(settings.Shortcuts.MouseMoveFastDown.Keys)
	vkCodesSetMousePosLeftFast := ts(settings.Shortcuts.MouseMoveFastLeft.Keys)
	vkCodesSetMousePosRightFast := ts(settings.Shortcuts.MouseMoveFastRight.Keys)
	vkCodesSetMousePosUpSlow := ts(settings.Shortcuts.MouseMoveSlowUp.Keys)
	vkCodesSetMousePosDownSlow := ts(settings.Shortcuts.MouseMoveSlowDown.Keys)
	vkCodesSetMousePosLeftSlow := ts(settings.Shortcuts.MouseMoveSlowLeft.Keys)
	vkCodesSetMousePosRightSlow := ts(settings.Shortcuts.MouseMoveSlowRight.Keys)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionDown, mousectl.SpeedFast), 0, vkCodesSetMousePosDownFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionUp, mousectl.SpeedFast), 0, vkCodesSetMousePosUpFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionLeft, mousectl.SpeedFast), 0, vkCodesSetMousePosLeftFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionRight, mousectl.SpeedFast), 0, vkCodesSetMousePosRightFast...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionUp, mousectl.SpeedSlow), 0, vkCodesSetMousePosUpSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionDown, mousectl.SpeedSlow), 0, vkCodesSetMousePosDownSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionLeft, mousectl.SpeedSlow), 0, vkCodesSetMousePosLeftSlow...)
	keyboardctl.RegisterOne(MoveMouseFunc(mousectl.DirectionRight, mousectl.SpeedSlow), 0, vkCodesSetMousePosRightSlow...)

	// I\R : in ModeControl, simulate mouse left button click
	// vkCodesMouseLeftClick := [][]uint32{{keyboardctl.VK_I}, {keyboardctl.VK_R}}
	vkCodesMouseLeftClick := [][]uint32{ts(settings.Shortcuts.MouseLeftButtonClickPrimary.Keys), ts(settings.Shortcuts.MouseLeftButtonClickSecondary.Keys)}
	keyboardctl.RegisterMulti(MouseLeftClick, 0, vkCodesMouseLeftClick...)

	// O\T : in ModeControl, simulate mouse right button click
	// vkCodesMouseRightClick := [][]uint32{{keyboardctl.VK_O}, {keyboardctl.VK_T}}
	vkCodesMouseRightClick := [][]uint32{ts(settings.Shortcuts.MouseRightButtonClickPrimary.Keys), ts(settings.Shortcuts.MouseRightButtonClickSecondary.Keys)}
	keyboardctl.RegisterMulti(MouseRightClick, 0, vkCodesMouseRightClick...)

	// C\N : in ModeControl, simulate mouse left button hold
	// vkCoodesLeftDown := [][]uint32{{keyboardctl.VK_C}, {keyboardctl.VK_N}}
	vkCoodesLeftDown := [][]uint32{ts(settings.Shortcuts.MouseLeftButtonHoldPrimary.Keys), ts(settings.Shortcuts.MouseLeftButtonHoldSecondary.Keys)}
	keyboardctl.RegisterWithReleaseEventMulti(MouseLeftDown, 0, vkCoodesLeftDown...)

	// in ModeControl ,control the speed of your mouse scroll
	vkCodesMulitiScrollSpeedLevel := [][]uint32{{keyboardctl.VK_LSHIFT, keyboardctl.VK_1}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_2}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_3}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_4}, {keyboardctl.VK_LSHIFT, keyboardctl.VK_5}}
	keyboardctl.RegisterMulti(ScrollSpeedLevelSwitch, 0, vkCodesMulitiScrollSpeedLevel...)
	// shift + H\J\K\L : in ModeControl, control the mouse scroll like vim
	// shift + W\A\S\D : in ModeControl, control the mouse scroll like fps game
	vkCodesMouseVerticalScrollDownFast := ts(settings.Shortcuts.MouseScrollFastDown.Keys)
	vkCodesMouseVerticalScrollUpFast := ts(settings.Shortcuts.MouseScrollFastUp.Keys)
	vkCodesMouseHorizontslScrollLeftFast := ts(settings.Shortcuts.MouseScrollFastLeft.Keys)
	vkCodesMouseHorizontslScrollRightFast := ts(settings.Shortcuts.MouseScrollFastRight.Keys)
	vkCodesMouseVerticalScrollDownSlow := ts(settings.Shortcuts.MouseScrollSlowDown.Keys)
	vkCodesMouseVerticalScrollUpSlow := ts(settings.Shortcuts.MouseScrollSlowUp.Keys)
	vkCodesMouseHorizontslScrollLeftSlow := ts(settings.Shortcuts.MouseScrollSlowLeft.Keys)
	vkCodesMouseHorizontslScrollRightSlow := ts(settings.Shortcuts.MouseScrollSlowRight.Keys)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalDown, mousectl.SpeedFast), 0, vkCodesMouseVerticalScrollDownFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalUp, mousectl.SpeedFast), 0, vkCodesMouseVerticalScrollUpFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalLeft, mousectl.SpeedFast), 0, vkCodesMouseHorizontslScrollLeftFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalRight, mousectl.SpeedFast), 0, vkCodesMouseHorizontslScrollRightFast...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalLeft, mousectl.SpeedSlow), 0, vkCodesMouseHorizontslScrollLeftSlow...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionHorizontalRight, mousectl.SpeedSlow), 0, vkCodesMouseHorizontslScrollRightSlow...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalDown, mousectl.SpeedSlow), 0, vkCodesMouseVerticalScrollDownSlow...)
	keyboardctl.RegisterOne(ScrollMouseFunc(mousectl.DirectionVerticalUp, mousectl.SpeedSlow), 0, vkCodesMouseVerticalScrollUpSlow...)

	// _____________________________________________________________________________________________________________________________11111111111111111111111111222222222222222222222222222222222223333333333333333
	// main keyboard event listener
	keyboardctl.RawKeyboardListener(keyboardctl.LowLevelKeyboardCallback)

}

func void() {

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
	logger.Infof("", "toggle mode to:%d", base.GetMode())
	ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	ui.ToggleTrayIcon()
}

func ToggleOptionView(wParam uintptr, vkCode, scanCode uint32) uintptr {
	if base.ToggleOptionViewState() {
		StartOptionView()
	} else {
		HideOptionView()
	}
	return 1
}

func StartControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d", base.GetMode())
	if base.GetMode() == base.ModeControl {
		logger.Infof("", "already in control mode:%s", time.Now())
	} else {
		base.SetMode(base.ModeControl)
		logger.Infof("", "change to control mode:%s", time.Now())
	}
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	return 1
}

func ResetSetting(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "user restore setting")
	config.RestoreSettings()
	return 1
}

func QuitControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d", base.GetMode())
	if base.GetMode() == base.ModeControl {
		logger.Infof("", "change to normal mode", time.Now())
		base.SetMode(base.ModeNormal)
	} else {
		logger.Infof("", "already in normal mode", time.Now())
	}
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	return 0
}

func ScrollSpeedLevelSwitch(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
	// if base.GetMode() != base.ModeControl {
	// 	logger.Infof("","not in control mode, can not switch speed,mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	if util.Contains[uint32](vkCodesMulitiSpeedLevelArr, uint32(vkCode)) {
		speedLevel := int(vkCode) - keyboardctl.VK_1 + 1
		base.SetScrollSpeedLevel(speedLevel)
		logger.Infof("", "change speed to :%d", base.GetScrollSpeed())
	}
	return 1
}

func SpeedLevelSwitch(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
	// if base.GetMode() != base.ModeControl {
	// 	logger.Infof("","not in control mode, can not switch speed,mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	if util.Contains[uint32](vkCodesMulitiSpeedLevelArr, uint32(vkCode)) {
		speedLevel := int(vkCode) - keyboardctl.VK_1 + 1
		base.SetMoveSpeedLevel(speedLevel)
		logger.Infof("", "change speed to :%d", base.GetMoveSpeedLevel())
	}
	return 1
}

func MouseLeftClick(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d", base.GetMode())
	// if base.GetMode() != base.ModeControl {
	// 	logger.Infof("","not in control mode, can not switch speed,mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	mousectl.LeftClick()
	logger.Infof("", "mouse left click")
	return 1
}

func MouseRightClick(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d", base.GetMode())
	// if base.GetMode() != base.ModeControl {
	// 	logger.Infof("","not in control mode, can not switch speed,mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	mousectl.RightClick()
	logger.Infof("", "mouse right click")
	return 1
}

func MouseLeftDown(wParam uintptr, vkCode, scanCode uint32) uintptr {
	if wParam == keyboardctl.WM_KEYDOWN {
		logger.Infof("", "mouse left button down")
		mousectl.MouseLeftDown()
	} else if wParam == keyboardctl.WM_KEYUP {
		logger.Infof("", "mouse left button up")
		mousectl.MouseLeftUp()
	}
	return 1
}

func MoveMouseFunc(direction mousectl.MoveDirection, speedType mousectl.MoveSpeedType) keyboardctl.Callback2 {
	return func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		// if base.GetMode() != base.ModeControl {
		// 	logger.Infof("","not in control mode, can not switch speed,mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel())
		// 	return 0
		// }
		mousectl.MoveMouseCtrl(direction, speedType)
		return 1
	}
}

func ScrollMouseFunc(direction mousectl.ScrollDirection, speed mousectl.MoveSpeedType) keyboardctl.Callback2 {
	return func(wParam uintptr, vkCode, scanCode uint32) uintptr { // if base.GetMode() != base.ModeControl { logger.Infof("","not in control mode, can not switch speed,mode:%d,current speed:%d", base.GetMode(), base.GetMoveSpeedLevel()) return 0 }
		mousectl.ScrollMouseCtrl(direction, speed)
		return 1
	}
}

func TmpQuitControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "current mode:%d", base.GetMode())
	if base.GetMode() == base.ModeControl {
		logger.Infof("", "change to normal mode", time.Now())
		base.SetMode(base.ModeNormal)
	} else {
		logger.Infof("", "already in normal mode", time.Now())
	}
	// ui.Message(fmt.Sprintf("change to: %s mode", base.GetModeDesc()))
	return 0
}

func ToggleControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	toggleControlMode()
	return 1
}

func ForceQuit(wParam uintptr, vkCode, scanCode uint32) uintptr {
	logger.Infof("", "force quit")
	os.Exit(0)
	return 1
}

func BlockKey(wParam uintptr, vkCode, scanCode uint32) uintptr {
	return 1
}

func ToggleHelpPane(wParam uintptr, vkCode, scanCode uint32) uintptr {
	toggleHelpPane()
	return 1
}

func toggleHelpPane() {
	base.SetHelperMode(1 - base.GetHelperMode())
	logger.Infof("", "toggle helper mode to:%d", base.GetMode())
	ui.ToggleHelper(fmt.Sprintf("change to: %s helper mode", base.GetModeDesc()), base.GetHelperMode())
}
