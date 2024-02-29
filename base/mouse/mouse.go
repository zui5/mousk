package mouse

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32       = windows.NewLazySystemDLL("user32.dll")
	sendInput    = user32.NewProc("SendInput")
	getCursorPos = user32.NewProc("GetCursorPos")
	setCursorPos = user32.NewProc("SetCursorPos")
)

const (
	inputMouse   = 0
	mouseMove    = 0x0001
	mouseEventFV = 0x8000
	absolute     = 0x8000
	inputMouseFV = 0
	move         = 1
	mouseEventF  = 0
)

type input struct {
	Type uint32
	Mi   mouseInput
}

type mouseInput struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

func MoveMouse(dx, dy int32) {
	var inputs [1]input
	inputs[0] = input{
		Type: inputMouse,
		Mi: mouseInput{
			Dx:      dx,
			Dy:      dy,
			DwFlags: mouseMove,
			Time:    0,
		},
	}
	sendInput.Call(
		uintptr(unsafe.Sizeof(input{})),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(unsafe.Sizeof(input{})),
	)
}

func GetMousePos() (int32, int32) {
	var pos struct {
		x, y int32
	}
	getCursorPos.Call(uintptr(unsafe.Pointer(&pos)))
	return pos.x, pos.y
}

func SetMousePos(x, y int32) {
	setCursorPos.Call(uintptr(x), uintptr(y))
}
