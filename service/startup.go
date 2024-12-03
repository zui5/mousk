package service

import (
	"mousk/common/logger"
	"mousk/infra/config"

	"github.com/wailsapp/wails/v3/plugins/experimental/start_at_login"
)

type StartupService struct{}

func (i *StartupService) Startup(startupstate bool) bool {
	logger.Infof("", "startup state:%t", startupstate)
	config.GetSettings().StartOnSystemUp = startupstate
	if err := config.WriteSettings(); err != nil {
		return false
	}
	start_at_login := start_at_login.NewPlugin(start_at_login.Config{
		RegistryKey: "mousk.exe",
	})
	if startupstate {
		start_at_login.StartAtLogin(true)
	} else {
		start_at_login.StartAtLogin(false)
	}
	return true
}

func (i *StartupService) StartupState() bool {
	return config.GetSettings().StartOnSystemUp
}
