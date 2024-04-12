package keyboardctl

import (
	"fmt"
	"mousek/infra/base"
	"syscall"
	"unsafe"
)

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
	WM_KEYUP       = 0x0101
	VK_SHIFT       = 0xa0
	VK_CONTROL     = 0xa2
	VK_A           = 0x41
	VK_Q           = 0x51
)

var (
	keyStates           map[uint32]bool
	setWindowsHookEx    = base.User32.NewProc("SetWindowsHookExW")
	getMessageW         = base.User32.NewProc("GetMessageW")
	unhookWindowsHookEx = base.User32.NewProc("UnhookWindowsHookEx")
)

type KBDLLHOOKSTRUCT struct {
	VkCode uint32
}

type Callback func(nCode int, wParam uintptr, lParam uintptr) uintptr
type HookProc func(nCode int, wParam uintptr, lParam uintptr) uintptr

func KeyboardListener(cb Callback) {

	keyStates = make(map[uint32]bool)

	hookProc := HookProc(cb)

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

func Pressed(vkCode uint32) bool {
	return keyStates[vkCode]
}

func SetPressed(vkCode uint32) {
	keyStates[vkCode] = true
}

func SetReleased(vkCode uint32) {
	keyStates[vkCode] = false
}
