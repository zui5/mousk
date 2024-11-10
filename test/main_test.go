package main

import (
	"log"
	"syscall"
	"testing"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	unhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	callNextHookEx      = user32.NewProc("CallNextHookEx")
	getMessageW         = user32.NewProc("GetMessageW")
	WH_KEYBOARD_LL      = 13
	WM_KEYDOWN          = 0x0100
	WM_KEYUP            = 0x0101
	keyPressedStates    map[uint32]bool
)

type KBDLLHOOKSTRUCT struct {
	vkCode      uint32
	scanCode    uint32
	flags       uint32
	time        uint32
	dwExtraInfo uintptr
}

type Callback func(keyCode uint32, pressed bool)

func HookProc(cb Callback) func(int, uintptr, uintptr) uintptr {
	return func(nCode int, wParam uintptr, lParam uintptr) uintptr {
		if nCode >= 0 {
			kbdllhookstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
			keyCode := kbdllhookstruct.vkCode

			var pressed bool
			if int(wParam) == WM_KEYDOWN {
				pressed = true
			} else if int(wParam) == WM_KEYUP {
				pressed = false
			}

			// 更新键状态
			keyPressedStates[keyCode] = pressed

			// 记录日志信息
			if pressed {
				log.Printf("Key pressed: %d\n", keyCode)
			} else {
				log.Printf("Key released: %d\n", keyCode)
			}

			// // 检查 Shift 键状态
			// if keyPressedStates[0x10] {
			// 	log.Println("Shift is held down")
			// } else {
			// 	log.Println("Shift is not held down")
			// }

			cb(keyCode, pressed)
		}

		ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
		return ret
	}
}

func RawKeyboardListener(cb Callback) {
	keyPressedStates = make(map[uint32]bool)
	hookProc := HookProc(cb)

	hookID, _, _ := setWindowsHookEx.Call(
		uintptr(WH_KEYBOARD_LL),
		syscall.NewCallback(hookProc),
		0,
		0,
	)

	if hookID == 0 {
		log.Println("Failed to set hook")
		return
	}

	log.Println("Hook set, waiting for events...")

	defer unhookWindowsHookEx.Call(hookID)

	var msg uintptr
	for {
		_, _, _ = getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}
}

func TestMain(t *testing.T) {

	RawKeyboardListener(func(keyCode uint32, pressed bool) {
		// 可以在这里添加额外的事件处理逻辑
	})
}
