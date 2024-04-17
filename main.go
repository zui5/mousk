package main

import (
	"fmt"
	"mousek/infra/keyboardctl"
	"mousek/infra/monitor"
	"mousek/infra/mousectl"
	"os"
	"time"
	"unsafe"
)

func main() {

	monitors := monitor.GetMonitors()
	for _, monitor := range monitors {
		moveMouseAround(monitor.Monitor)
	}

	keyboardctl.RawKeyboardListener(Callback)

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
		mousectl.MoveMouse(x, y)
		x += 10
		time.Sleep(5 * time.Millisecond)
	}

	// 向下移动到显示器底边缘
	for y < int(monitor.Bottom) {
		mousectl.MoveMouse(x, y)
		y += 10
		time.Sleep(5 * time.Millisecond)
	}

	// 向左移动到显示器左边缘
	for x > int(monitor.Left) {
		mousectl.MoveMouse(x, y)
		x -= 10
		time.Sleep(5 * time.Millisecond)
	}

	// 向上移动到显示器上边缘
	for y > int(monitor.Top) {
		mousectl.MoveMouse(x, y)
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
		if keyboardctl.Pressed(keyboardctl.VK_CONTROL) && keyboardctl.Pressed(keyboardctl.VK_SHIFT) && keyboardctl.Pressed(keyboardctl.VK_A) {
			fmt.Println("Ctrl+Shift+A keys pressed simultaneously")
		}

		// 如果按下了 'Q' 键，退出程序
		if keyboardctl.Pressed(keyboardctl.VK_Q) {
			os.Exit(0)
		}
		return 1
	}
	return 0
}
