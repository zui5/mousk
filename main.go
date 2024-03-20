package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
	VK_SHIFT       = 0xA0
	VK_CONTROL     = 0xA2
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

func main() {
	for {
		time.Sleep(time.Second)
		fmt.Println(isKeyDown(VK_SHIFT) && isKeyDown(VK_CONTROL))

	}
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

func Callback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 && wParam == WM_KEYDOWN {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vkCode := kbdStruct.VkCode
		fmt.Printf("Key pressed (VK code): %x\n", vkCode)

		fmt.Println(isKeyDown(VK_SHIFT))

		// 如果同时按下了Shift和Control键，则打印消息
		// if isKeyDown(VK_SHIFT) || isKeyDown(VK_CONTROL) {
		// fmt.Println("Shift and Control keys pressed simultaneously")
		// }
		return 1
	}

	return 0
}

// 辅助函数，检查指定的虚拟键是否被按下
func isKeyDown(vkCode uint32) bool {
	a := GetAsyncKeyState(vkCode) & 0x8000
	return a != 0
}

// 获取指定虚拟键的当前状态
func GetAsyncKeyState(vkCode uint32) int32 {
	user32 := syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState := user32.NewProc("GetAsyncKeyState")
	ret, _, _ := getAsyncKeyState.Call(uintptr(vkCode))
	return int32(ret)
}
