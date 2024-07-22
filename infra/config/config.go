package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var settingsVar Settings

var filePath = "../../confs.toml"

// var filePath = "./confs.toml"

const defaultSettings = `
StartOnSystemUp = true

[PresetFunc]
ChangeMode =["space","tab"]
OpenSetting = ["space",","]
  
[PresetFunc.MouseMove]
  [PresetFunc.MouseMove.Fast]
  Down =["j"]
  Left =["h"]
  Right =["l"]
  Up =["k"]
  [PresetFunc.MouseMove.Slow]
  Down =["s"]
  Left =["a"]
  Right =["d"]
  Up =["w"]
  [PresetFunc.MouseMove.SpeedLevel]
  Level1 = 1
  Level2 = 2
  Level3 = 3
  Level4 = 4
  Level5 = 5

[PresetFunc.MouseScroll]
  [PresetFunc.MouseScroll.Fast]
  Down =["shift","j"]
  Left =["shift","h"]
  Right =["shift","l"]
  Up =["shift","k"]
  [PresetFunc.MouseScroll.Slow]
  Down =["shift","s"]
  Left =["shift","a"]
  Right =["shift","d"]
  Up =["shift","w"]
  [PresetFunc.MouseScroll.SpeedLevel]
  Level1 = 1
  Level2 = 2
  Level3 = 3
  Level4 = 4
  Level5 = 5

[PresetFunc.MouseLeftButtonClick]
  Primary = "i"
  Secondary = ""
[PresetFunc.MouseRightButtonClick]
  Primary = "o"
  Secondary = ""
[PresetFunc.MouseLeftButtonHold]
  Primary = "c"
  Secondary = ""
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

// type Settings struct {
// 	Mouse struct {
// 		Fast MouseConfig
// 		Slow MouseConfig
// 	}
// }

type Settings struct {
	StartOnSystemUp bool `toml:"StartOnSystemUp"`
	PresetFunc      struct {
		ChangeMode  []string `toml:"ChangeMode"`
		OpenSetting []string `toml:"OpenSetting"`
		MouseMove   struct {
			Fast struct {
				Down  []string `toml:"Down"`
				Left  []string `toml:"Left"`
				Right []string `toml:"Right"`
				Up    []string `toml:"Up"`
			} `toml:"Fast"`
			Slow struct {
				Down  []string `toml:"Down"`
				Left  []string `toml:"Left"`
				Right []string `toml:"Right"`
				Up    []string `toml:"Up"`
			} `toml:"Slow"`
			SpeedLevel struct {
				Level1 int `toml:"Level1"`
				Level2 int `toml:"Level2"`
				Level3 int `toml:"Level3"`
				Level4 int `toml:"Level4"`
				Level5 int `toml:"Level5"`
			} `toml:"SpeedLevel"`
		} `toml:"MouseMove"`
		MouseScroll struct {
			Fast struct {
				Down  []string `toml:"Down"`
				Left  []string `toml:"Left"`
				Right []string `toml:"Right"`
				Up    []string `toml:"Up"`
			} `toml:"Fast"`
			Slow struct {
				Down  []string `toml:"Down"`
				Left  []string `toml:"Left"`
				Right []string `toml:"Right"`
				Up    []string `toml:"Up"`
			} `toml:"Slow"`
			SpeedLevel struct {
				Level1 int `toml:"Level1"`
				Level2 int `toml:"Level2"`
				Level3 int `toml:"Level3"`
				Level4 int `toml:"Level4"`
				Level5 int `toml:"Level5"`
			} `toml:"SpeedLevel"`
		} `toml:"MouseScroll"`
		MouseLeftButtonClick struct {
			Primary   string `toml:"Primary"`
			Secondary string `toml:"Secondary"`
		} `toml:"MouseLeftButtonClick"`
		MouseRightButtonClick struct {
			Primary   string `toml:"Primary"`
			Secondary string `toml:"Secondary"`
		} `toml:"MouseRightButtonClick"`
		MouseLeftButtonHold struct {
			Primary   string `toml:"Primary"`
			Secondary string `toml:"Secondary"`
		} `toml:"MouseLeftButtonHold"`
	} `toml:"PresetFunc"`
}

func Init() {
	if InitConfigFile() != nil {
		log.Fatalf("init config file error")
		return
	}

	if LoadSettingsFromFile() != nil {
		log.Fatalf("load config file error")
		return
	}
}

func LoadSettingsFromFile() error {
	_, filePath, err := getConfigPath()
	if err != nil {
		return err
	}
	if _, err := toml.DecodeFile(filePath, &settingsVar); err != nil {
		return fmt.Errorf("decoding TOML:%+v", err)
	}

	// 输出生成的结构体
	fmt.Printf("%#v\n", settingsVar)
	return nil
}

func GetSettings() *Settings {
	return &settingsVar
}

// TODO
func ChangeSettings() {

}

func WriteSettings() error {
	_, filePath, err := getConfigPath()
	if err != nil {
		return err
	}
	// 打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open TOML:%+v", err)
	}
	defer file.Close()
	// 编码为TOML格式并写入文件
	if err := toml.NewEncoder(file).Encode(settingsVar); err != nil {
		return fmt.Errorf("writeback TOML:%+v", err)
	}
	return nil
}

func getConfigPath() (string, string, error) {
	// 获取当前用户的主目录
	usr, err := user.Current()
	if err != nil {
		return "", "", err
	}

	// 构建文件路径
	configDir := filepath.Join(usr.HomeDir, ".config", "mousek")
	filePath := filepath.Join(configDir, "confs.toml")
	return configDir, filePath, nil
}

func InitConfigFile() error {
	// 构建文件路径
	configDir, filePath, err := getConfigPath()
	if err != nil {
		return err
	}
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 创建目录
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			return fmt.Errorf("creating directories:%+v", err)
		}

		// 创建文件
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("creating file:%+v", err)
		}
		defer file.Close()

		err = RestoreSettings()
		if err != nil {
			return err
		}
		fmt.Println("Config file created at:", filePath)
	} else {
		err = LoadSettingsFromFile()
		if err != nil {
			return err
		}
		fmt.Println("Config file already exists at:", filePath)
	}
	return nil
}

func RestoreSettings() error {
	if _, err := toml.Decode(defaultSettings, &settingsVar); err != nil {
		return fmt.Errorf("decoding TOML:%+v", err)
	}
	// 输出生成的结构体
	fmt.Printf("%#v\n", settingsVar)
	if err := WriteSettings(); err != nil {
		return err
	}
	return nil
}
