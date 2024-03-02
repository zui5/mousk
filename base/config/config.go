package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var settingsVar Settings

// var filePath = "../../confs.toml"
var filePath = "./confs.toml"

const defaultSettings = `
[Mouse]
  [Mouse.Slow]
    Speed = 4
    Up = "k"
    Down = "j"
    Left = "h"
    Right = "l"
  [Mouse.Fast]
    Speed = 20
    Up = "w"
    Down = "s"
    Left = "a"
    Right = "d"
`

// 配置项枚举
const (
	MouseSlowSpeed = iota
	MouseSlowUp
	MouseSlowDown
	MouseSlowLeft
	MouseSlowRight
	MouseFastSpeed
	MouseFastUp
	MouseFastDown
	MouseFastLeft
	MouseFastRight

	KeyDown = iota
	KeyHold
	KeyUp

	MouseUp
	MouseHold
	MouseDown
	MouseMove
	MouseDrag
	MouseWheel

	WheelUp
	WheelDown
)

type MouseConfig struct {
	Speed int
	Up    string
	Down  string
	Left  string
	Right string
}

type Settings struct {
	Mouse struct {
		Fast MouseConfig
		Slow MouseConfig
	}
}

func Init() {
	LoadSettingsFromFile()
}

func LoadSettingsFromFile() {
	if _, err := toml.DecodeFile(filePath, &settingsVar); err != nil {
		fmt.Println("Error decoding TOML:", err)
		return
	}
	// 输出生成的结构体
	fmt.Printf("%#v\n", settingsVar)
}

func GetSettings() *Settings {
	return &settingsVar
}

// TODO
func ChangeSettings() {

}

func WriteSettings() {
	// 打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error Open TOML:", err)
		return
	}
	defer file.Close()

	// 编码为TOML格式并写入文件
	if err := toml.NewEncoder(file).Encode(settingsVar); err != nil {
		fmt.Println("Error Writeback TOML:", err)
		return
	}
}

func RestoreSettings() {
	if _, err := toml.Decode(defaultSettings, &settingsVar); err != nil {
		fmt.Println("Error decoding TOML:", err)
		return
	}
	// 输出生成的结构体
	fmt.Printf("%#v\n", settingsVar)
	WriteSettings()
}
