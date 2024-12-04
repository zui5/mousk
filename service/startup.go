package service

import (
	"mousk/common/logger"
	"mousk/infra/config"
	"os"
	"path/filepath"

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
		RegistryKey: getExcutable(),
	})

	if startupstate {
		start_at_login.StartAtLogin(true)
	} else {
		start_at_login.StartAtLogin(false)
	}
	return true
}

func getExcutable() string {
	// 获取当前可执行文件的目录
	exePath, err := os.Executable()
	if err != nil {
		return "mousk.exe"
	}
	exeDir := filepath.Dir(exePath)

	return exeDir
}

func (i *StartupService) StartupState() bool {
	return config.GetSettings().StartOnSystemUp
}
