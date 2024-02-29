package monitor

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32              = windows.NewLazySystemDLL("user32.dll")
	enumDisplayMonitors = user32.NewProc("EnumDisplayMonitors")
	getMonitorInfo      = user32.NewProc("GetMonitorInfoW")
)

// MONITORINFOEX structure contains information about a display monitor.
type MONITORINFOEX struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
	SzDevice  [32]uint16
}

// RECT structure defines the coordinates of the upper-left and lower-right corners of a rectangle.
type RECT struct {
	Left, Top, Right, Bottom int32
}

// MonitorCallback is a callback function used with EnumDisplayMonitors.
func MonitorCallback(hMonitor, hdcMonitor, lprcMonitor uintptr, dwData uintptr) uintptr {
	// Get additional information using GetMonitorInfoW
	var monitorInfoW MONITORINFOEX
	monitorInfoW.CbSize = uint32(unsafe.Sizeof(monitorInfoW))
	getMonitorInfo.Call(hMonitor, uintptr(unsafe.Pointer(&monitorInfoW)))

	// Convert the array to a slice before passing it to UTF16ToString.
	deviceName := windows.UTF16ToString(monitorInfoW.SzDevice[:])

	fmt.Printf("Monitor: %s\n", deviceName)
	fmt.Printf("  Monitor Rectangle: (%d, %d, %d, %d)\n", monitorInfoW.RcMonitor.Left, monitorInfoW.RcMonitor.Top, monitorInfoW.RcMonitor.Right, monitorInfoW.RcMonitor.Bottom)
	fmt.Printf("  Work Area Rectangle: (%d, %d, %d, %d)\n", monitorInfoW.RcWork.Left, monitorInfoW.RcWork.Top, monitorInfoW.RcWork.Right, monitorInfoW.RcWork.Bottom)
	fmt.Println()
	return 1
}

func EnumDisplay() {
	// Enumerate display monitors
	enumDisplayMonitors.Call(0, 0, syscall.NewCallback(MonitorCallback), 0)
}
