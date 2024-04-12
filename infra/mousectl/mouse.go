package mousectl

import (
	"mousek/infra/base"
)

var (
	setCursorPos = base.User32.NewProc("SetCursorPos")
)

// 设置鼠标位置
func MoveMouse(x, y int) {
	setCursorPos.Call(uintptr(x), uintptr(y))
}
