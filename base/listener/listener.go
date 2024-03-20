package listener

import (
	"mousek/base/callback"
	"mousek/base/config"
	"mousek/base/mouse"

	hook "mousek/base/hook"
)

// var RegisterKeys []RegisterKey
var registerFuncList []RegisterFunc

type RegisterFunc struct {
	Key          RegisterKey
	CallbackFunc func(hook.Event)
}
type RegisterKey struct {
	Cmd  []string
	When uint8
}

// add one listner func to var registerFncList
func RegisterOne(key RegisterKey, cb callback.CallbackFunc) {
	registerFuncList = append(registerFuncList, RegisterFunc{key, cb})
}

// ergodic the registerFuncList register event and start
func Start() {
	for _, v := range registerFuncList {
		hook.Register(v.Key.When, v.Key.Cmd, v.CallbackFunc)
	}

	s := hook.Start()
	<-hook.Process(s)
}

func RegisterFromConfig() {
	settingsVar := config.GetSettings()
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Fast.Down}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedFast, mouse.DirectionDown))
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Fast.Up}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedFast, mouse.DirectionUp))
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Fast.Left}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedFast, mouse.DirectionLeft))
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Fast.Right}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedFast, mouse.DirectionRight))

	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Slow.Down}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedSlow, mouse.DirectionDown))
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Slow.Up}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedSlow, mouse.DirectionUp))
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Slow.Left}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedSlow, mouse.DirectionLeft))
	RegisterOne(RegisterKey{Cmd: []string{settingsVar.Mouse.Slow.Right}, When: hook.KeyDown}, callback.MouseMove(mouse.SpeedSlow, mouse.DirectionRight))
}
