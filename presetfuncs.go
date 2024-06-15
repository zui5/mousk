package main

import (
	"fmt"
	"mousek/infra/base"
	"mousek/infra/keyboardctl"
	"mousek/infra/mousectl"
	"mousek/infra/util"
	"time"
)

func ToggleControlMode() {
	base.SetMode(1 - base.GetMode())
	fmt.Printf("toggle mode to:%d\n", base.GetMode())
}

func StartControlMode(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d\n", base.GetMode())
	if base.GetMode() == base.ModeControl {
		fmt.Println("already in control mode", time.Now())
	} else {
		base.SetMode(base.ModeControl)
		fmt.Println("change to control mode", time.Now())
	}
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
	return 0
}

func ScrollSpeedLevelSwitch(wParam uintptr, vkCode, scanCode uint32) uintptr {
	fmt.Printf("current mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// if base.GetMode() != base.ModeControl {
	// 	fmt.Printf("not in control mode, can not switch speed,mode:%d,current speed:%d\n", base.GetMode(), base.GetMoveSpeedLevel())
	// 	return 0
	// }
	if util.Contains[uint32](vkCodesMulitiSpeedLevelArr, uint32(vkCode)) {
		speedLevel := int(vkCode) - keyboardctl.VK_0 + 1
		base.SetScrollSpeedLevel(speedLevel)
		fmt.Printf("change speed to :%d\n", base.GetScrollSpeedLevel())
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
		speedLevel := int(vkCode) - keyboardctl.VK_0 + 1
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
