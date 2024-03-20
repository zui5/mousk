package syscallwrapper

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
)

type KBDLLHOOKSTRUCT struct {
	VkCode uint32
}

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	getMessageW         = user32.NewProc("GetMessageW")
	callNextHookEx      = user32.NewProc("CallNextHookEx")
	unhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
)

type HookProc func(nCode int, wParam uintptr, lParam uintptr) uintptr

func main() {
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

	var msg uintptr
	for {
		_, _, _ = getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}

	defer unhookWindowsHookEx.Call(hookID)
}

func Callback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 && wParam == WM_KEYDOWN {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vkCode := kbdStruct.VkCode
		fmt.Printf("Key pressed (VK code): %d\n", vkCode)
	}

	ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
	return ret
}
