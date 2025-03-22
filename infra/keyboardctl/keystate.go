package keyboardctl

import (
	"mousk/infra/base"
)

var (
	procGetAsyncKeyState = base.User32.NewProc("GetAsyncKeyState")
)

func IsShiftPressed() bool {
	v, _, _ := procGetAsyncKeyState.Call(uintptr(VK_SHIFT))
	return v != 0
}
