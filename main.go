package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var (
	enumDisplayMonitors = user32.NewProc("EnumDisplayMonitors")
)
var (
	setCursorPos     = user32.NewProc("SetCursorPos")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

type MONITORINFOEX struct {
	Size    uint32
	Monitor RECT
	Work    RECT
	Flags   uint32
}

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
	WM_KEYUP       = 0x0101
	VK_SHIFT       = 0xa0
	VK_CONTROL     = 0xa2
	VK_A           = 0x41
	VK_Q           = 0x51
)

type KBDLLHOOKSTRUCT struct {
	VkCode uint32
}

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	getMessageW         = user32.NewProc("GetMessageW")
	unhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
)

type HookProc func(nCode int, wParam uintptr, lParam uintptr) uintptr

var (
	keyStates map[uint32]bool
)

func main() {

	// 显示器信息获取
	var monitors []MONITORINFOEX

	// 定义回调函数
	callback := syscall.NewCallback(func(hMonitor uintptr, hdc uintptr, lprc uintptr, dwData uintptr) uintptr {
		var mi MONITORINFOEX
		mi.Size = uint32(unsafe.Sizeof(mi))
		if ret, _, _ := enumDisplayMonitors.Call(hdc, uintptr(unsafe.Pointer(&RECT{})), syscall.NewCallback(func(lprc uintptr, hdc uintptr, lprcClip uintptr, dwData uintptr) uintptr {
			return 1
		}), 0); ret != 0 {
			if ret, _, _ := syscall.Syscall(user32.NewProc("GetMonitorInfoW").Addr(), 2, hMonitor, uintptr(unsafe.Pointer(&mi)), 0); ret != 0 {
				monitors = append(monitors, mi)
			}
		}
		return 1
	})

	// 调用 EnumDisplayMonitors 函数
	if ret, _, _ := enumDisplayMonitors.Call(0, 0, callback, 0); ret == 0 {
		fmt.Println("EnumDisplayMonitors failed")
		return
	}

	// 输出显示器数量和范围信息
	fmt.Printf("Number of monitors: %d\n", len(monitors))
	for i, monitor := range monitors {
		fmt.Printf("Monitor %d:\n", i+1)
		fmt.Printf("    Left:   %d\n", monitor.Monitor.Left)
		fmt.Printf("    Top:    %d\n", monitor.Monitor.Top)
		fmt.Printf("    Right:  %d\n", monitor.Monitor.Right)
		fmt.Printf("    Bottom: %d\n", monitor.Monitor.Bottom)
		moveMouseAround(monitor.Monitor)
	}

	//

	// 键盘事件监听
	keyStates = make(map[uint32]bool)

	hookProc := HookProc(Callback)

	hookID, _, _ := setWindowsHookEx.Call(
		uintptr(WH_KEYBOARD_LL),
		syscall.NewCallback(hookProc),
		0,
		0,
	)

	if hookID == 0 {
		fmt.Println("Failed to set hook")
		return
	}

	fmt.Println("Hook set, waiting for events...")

	defer unhookWindowsHookEx.Call(hookID)

	var msg uintptr
	for {
		_, _, _ = getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}
}

// 控制鼠标在指定显示器的四周移动
func moveMouseAround(monitor RECT) {
	x := int(monitor.Left)
	y := int(monitor.Top)

	/* 	width := int(monitor.Right - monitor.Left)
	   	height := int(monitor.Bottom - monitor.Top)
	*/
	// 向右移动到显示器右边缘
	for x < int(monitor.Right) {
		moveMouse(x, y)
		x++
		time.Sleep(5 * time.Millisecond)
	}

	// 向下移动到显示器底边缘
	for y < int(monitor.Bottom) {
		moveMouse(x, y)
		y++
		time.Sleep(5 * time.Millisecond)
	}

	// 向左移动到显示器左边缘
	for x > int(monitor.Left) {
		moveMouse(x, y)
		x--
		time.Sleep(5 * time.Millisecond)
	}

	// 向上移动到显示器上边缘
	for y > int(monitor.Top) {
		moveMouse(x, y)
		y--
		time.Sleep(5 * time.Millisecond)
	}
}

// 设置鼠标位置
func moveMouse(x, y int) {
	setCursorPos.Call(uintptr(x), uintptr(y))
}
func Callback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vkCode := kbdStruct.VkCode

		if wParam == WM_KEYDOWN {
			keyStates[vkCode] = true
			fmt.Printf("Key pressed (VK code): %x\n", vkCode)
		} else if wParam == WM_KEYUP {
			keyStates[vkCode] = false
			fmt.Printf("Key released (VK code): %x\n", vkCode)
		}

		// 检查是否同时按下了 Ctrl、Shift 和 A 键
		if keyStates[VK_CONTROL] && keyStates[VK_SHIFT] && keyStates[VK_A] {
			fmt.Println("Ctrl+Shift+A keys pressed simultaneously")
		}

		// 如果按下了 'Q' 键，退出程序
		if keyStates[VK_Q] {
			os.Exit(0)
		}

		return 1
	}

	return 0
}
