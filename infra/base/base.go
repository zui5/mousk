package base

import "syscall"

var User32 = syscall.NewLazyDLL("user32.dll")
