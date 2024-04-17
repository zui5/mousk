package keyboardctl

import (
	"fmt"
	"mousek/infra/base"
	"os"
	"syscall"
	"unsafe"
)

var (
	keyPressedStates    map[uint32]bool
	setWindowsHookEx    = base.User32.NewProc("SetWindowsHookExW")
	getMessageW         = base.User32.NewProc("GetMessageW")
	unhookWindowsHookEx = base.User32.NewProc("UnhookWindowsHookEx")
	// todo serilize to file config
	listeningKeyReference = make(map[uint32]*KeyReference, 0)
)

type KBDLLHOOKSTRUCT struct {
	VkCode uint32
}
type KeyReference struct {
	Count           int
	KeyCombinations []KeyCallback
}

type KeyCallback struct {
	Keys []uint32 // for example. []{"ctrl","a"}
	Cb   Callback2
}

type Callback HookProc
type Callback2 HookProc
type HookProc func(nCode int, wParam uintptr, lParam uintptr) uintptr

func registerKeyListening(cb Callback2, vkCodes ...uint32) {
	for _, v := range vkCodes {
		listeningKeyReference[v].Count += 1
		listeningKeyReference[v].KeyCombinations = append(listeningKeyReference[v].KeyCombinations, KeyCallback{
			Keys: vkCodes,
			Cb:   cb,
		})
	}
}

func unRegisterKeyListening(vkCodes ...uint32) {
	for _, v := range vkCodes {
		listeningKeyReference[v].Count -= 1
		listeningKeyReference[v].KeyCombinations = append(listeningKeyReference[v].KeyCombinations, KeyCallback{
			Keys: vkCodes,
		})
	}

}
func Register(keyAction int, cb Callback, vkCodes ...uint32) {
	switch keyAction {
	case WM_KEYDOWN:

		break
	default:
		fmt.Println("method not develop yet")
		return
	}
	registerKeyListening(vkCodes...)
}

func Callback1(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vkCode := kbdStruct.VkCode

		if wParam == WM_KEYDOWN {
			SetPressed(vkCode)
			fmt.Printf("Key pressed (VK code): %x\n", vkCode)
		} else if wParam == WM_KEYUP {
			SetReleased(vkCode)
			fmt.Printf("Key released (VK code): %x\n", vkCode)
		}

		// 检查是否同时按下了 Ctrl、Shift 和 A 键
		if Pressed(VK_CONTROL) && Pressed(VK_SHIFT) && Pressed(VK_A) {
			fmt.Println("Ctrl+Shift+A keys pressed simultaneously")
		}

		// 如果按下了 'Q' 键，退出程序
		if Pressed(VK_Q) {
			os.Exit(0)
		}
		return 1
	}
	return 0
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
	return keyPressedStates[vkCode]
}

func AllPressed(vkCodes ...uint32) bool {
	if vkCodes == nil {
		return true
	}
	allPressed := true
	for _, v := range vkCodes {
		allPressed = allPressed && Pressed(v)
	}
	return allPressed
}

func SetPressed(vkCode uint32) {
	keyPressedStates[vkCode] = true
}

func SetReleased(vkCode uint32) {
	keyPressedStates[vkCode] = false
}
