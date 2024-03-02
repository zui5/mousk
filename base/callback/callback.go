package callback

import (
	"changeme/base/config"
	"changeme/base/mouse"

	hook "github.com/robotn/gohook"
)

type CallbackFunc func(hook.Event)

func MouseMove(speedLevel, direction int) CallbackFunc {
	return func(e hook.Event) {
		settingsVar := config.GetSettings()
		speed := settingsVar.Mouse.Slow.Speed
		dx := 0
		dy := 0
		if speedLevel == mouse.SpeedFast {
			speed = settingsVar.Mouse.Fast.Speed
		}
		switch direction {
		case mouse.DirectionUp:
			dx, dy = 0, -1*speed
		case mouse.DirectionDown:
			dx, dy = 0, 1*speed
		case mouse.DirectionLeft:
			dx, dy = -1*speed, 0
		case mouse.DirectionRight:
			dx, dy = 1*speed, 0
		default:
		}
		mouse.MoveMouse(int32(dx), int32(dy))
	}
}
