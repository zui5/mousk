package main

import (
	"fmt"
	"mousek/infra/base"
	"mousek/infra/keyboardctl"
	"mousek/infra/monitor"
	"mousek/infra/mousectl"
	"mousek/infra/util"
	"os"
	"time"
	"unsafe"
)

var vkCodesMulitiSpeedLevelArr = []uint32{keyboardctl.VK_1, keyboardctl.VK_2, keyboardctl.VK_3, keyboardctl.VK_4, keyboardctl.VK_5}

const (
	ModeNormal  = 0
	ModeControl = 1
)

func main() {

	// monitors := monitor.GetMonitors()
	// for _, monitor := range monitors {
	// 	moveMouseAround(monitor.Monitor)
	// }

	// win+space : activate control mode
	vkCodesWinSpace := []uint32{keyboardctl.VK_LWIN, keyboardctl.VK_SPACE}
	startControlMode := func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		fmt.Printf("current mode:%d\n", base.GetMode())
		fmt.Println()
		if base.GetMode() == ModeControl {
			fmt.Println("already in control mode", time.Now())
		} else {
			base.SetMode(ModeControl)
			fmt.Println("change to control mode", time.Now())
		}
		return 0
	}
	keyboardctl.RegisterOne(startControlMode, vkCodesWinSpace...)

	// when in ModeControl, 1\2\3\4...,control the speed of your mouse move
	vkCodesMulitiSpeedLevel := [][]uint32{{keyboardctl.VK_1}, {keyboardctl.VK_2}, {keyboardctl.VK_3}, {keyboardctl.VK_4}, {keyboardctl.VK_5}}
	speedLevelSwitch := func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		fmt.Printf("current mode:%d,current speed:%d\n", base.GetMode(), base.GetSpeedLevel())
		if base.GetMode() != ModeControl {
			fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetSpeedLevel())
			return 0
		}
		if util.Contains[uint32](vkCodesMulitiSpeedLevelArr, uint32(vkCode)) {
			speedLevel := int(vkCode) - keyboardctl.VK_0 + 1
			base.SetSpeedLevel(speedLevel)
			fmt.Printf("change speed to :%d\n", base.GetSpeedLevel())
		}
		return 0
	}
	keyboardctl.RegisterMulti(speedLevelSwitch, vkCodesMulitiSpeedLevel...)

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

	vkCodesMouseLeftClick := [][]uint32{{keyboardctl.VK_I}, {keyboardctl.VK_R}}
	mouseLeftClick := func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		fmt.Printf("current mode:%d\n", base.GetMode())
		if base.GetMode() != ModeControl {
			fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetSpeedLevel())
			return 0
		}
		mousectl.LeftClick()
		fmt.Printf("mouse left click\n")
		return 0
	}
	keyboardctl.RegisterMulti(mouseLeftClick, vkCodesMouseLeftClick...)

	vkCodesMouseRightClick := [][]uint32{{keyboardctl.VK_O}, {keyboardctl.VK_T}}
	mouseRightClick := func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		fmt.Printf("current mode:%d\n", base.GetMode())
		if base.GetMode() != ModeControl {
			fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetSpeedLevel())
			return 0
		}
		mousectl.RightClick()
		fmt.Printf("mouse right click\n")
		return 0
	}
	keyboardctl.RegisterMulti(mouseRightClick, vkCodesMouseRightClick...)

	mouseRightClickPress := func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		fmt.Printf("current mode:%d\n", base.GetMode())
		if base.GetMode() != ModeControl {
			fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetSpeedLevel())
			return 0
		}
		fmt.Printf("mouse left pressed\n")
		mousectl.RightClickLongPress(1 * time.Second)
		time.Sleep(1 * time.Second)
		return 0
	}
	keyboardctl.RegisterMulti(mouseRightClickPress, vkCodesMouseRightClickPress...)

	keyboardctl.RawKeyboardListener(keyboardctl.LowLevelKeyboardCallback)

}
func MoveMouseFunc(direction mousectl.MoveDirection, speedType mousectl.MoveSpeedType) keyboardctl.Callback2 {
	return func(wParam uintptr, vkCode, scanCode uint32) uintptr {
		if base.GetMode() != ModeControl {
			fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetSpeedLevel())
			return 0
		}
		mousectl.MoveMouseCtrl(direction, speedType)
		return 0
	}
}

// 控制鼠标在指定显示器的四周移动
func moveMouseAround(monitor monitor.RECT) {
	x := int(monitor.Left)
	y := int(monitor.Top)

	/* 	width := int(monitor.Right - monitor.Left)
	   	height := int(monitor.Bottom - monitor.Top)
	*/
	// 向右移动到显示器右边缘
	for x < int(monitor.Right) {
		mousectl.SetMousePos(x, y)
		x += 10
		time.Sleep(5 * time.Millisecond)
	}

	// 向下移动到显示器底边缘
	for y < int(monitor.Bottom) {
		mousectl.SetMousePos(x, y)
		y += 10
		time.Sleep(5 * time.Millisecond)
	}

	// 向左移动到显示器左边缘
	for x > int(monitor.Left) {
		mousectl.SetMousePos(x, y)
		x -= 10
		time.Sleep(5 * time.Millisecond)
	}

	// 向上移动到显示器上边缘
	for y > int(monitor.Top) {
		mousectl.SetMousePos(x, y)
		y -= 10
		time.Sleep(5 * time.Millisecond)
	}
}

func Callback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 {
		kbdStruct := (*keyboardctl.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vkCode := kbdStruct.VkCode

		if wParam == keyboardctl.WM_KEYDOWN {
			keyboardctl.SetPressed(vkCode)
			fmt.Printf("Key pressed (VK code): %x\n", vkCode)
		} else if wParam == keyboardctl.WM_KEYUP {
			keyboardctl.SetReleased(vkCode)
			fmt.Printf("Key released (VK code): %x\n", vkCode)
		}

		// 检查是否同时按下了 Ctrl、Shift 和 A 键
		// if keyboardctl.Pressed(keyboardctl._VK_CTRL) && keyboardctl.Pressed(keyboardctl._VK_SHIFT) && keyboardctl.Pressed(keyboardctl.VK_A) {
		// 	fmt.Println("Ctrl+Shift+A keys pressed simultaneously")
		// }

		// 如果按下了 'Q' 键，退出程序
		if keyboardctl.Pressed(keyboardctl.VK_Q) {
			os.Exit(0)
		}
		return 1
	}
	return 0
}
