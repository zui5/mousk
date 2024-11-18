package config

import "testing"

func TestInit(t *testing.T) {
	Init()
	// settingsVar.Mouse.Fast.Speed = 1000
	// settingsVar.PresetFunc.MouseMove.SpeedLevel.Level1 = -1
	WriteSettings()
	RestoreSettings()
}
