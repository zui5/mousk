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
  ActiveMode = ["SPACE", "TAB"]
  QuitMode = ["SPACE", "ESC"]
  OpenSetting = ["SPACE", "COMMA"]
  [PresetFunc.MouseMove]
    [PresetFunc.MouseMove.Fast]
      Down = ["J"]
      Left = ["H"]
      Right = ["L"]
      Up = ["K"]
    [PresetFunc.MouseMove.Slow]
      Down = ["S"]
      Left = ["A"]
      Right = ["D"]
      Up = ["W"]
    [PresetFunc.MouseMove.SpeedLevel]
      Level1 = 1
      LevelSwitch = [["1","2","3","4","5"]]
      Level2 = 2
      Level3 = 3
      Level4 = 4
      Level5 = 5
  [PresetFunc.MouseScroll]
    [PresetFunc.MouseScroll.Fast]
      Down = ["SHIFT", "J"]
      Left = ["SHIFT", "H"]
      Right = ["SHIFT", "L"]
      Up = ["SHIFT", "K"]
    [PresetFunc.MouseScroll.Slow]
      Down = ["SHIFT", "S"]
      Left = ["SHIFT", "A"]
      Right = ["SHIFT", "D"]
      Up = ["SHIFT", "W"]
    [PresetFunc.MouseScroll.SpeedLevel]
      LevelSwitch = [["SHIFT","1"],["SHIFT","2"],["SHIFT","3"],["SHIFT","4"],["SHIFT","5"]]
      Level1 = 1
      Level2 = 2
      Level3 = 3
      Level4 = 4
      Level5 = 5
  [PresetFunc.MouseLeftButtonClick]
    Primary = "I"
    Secondary = "R"
  [PresetFunc.MouseRightButtonClick]
    Primary = "O"
    Secondary = "T"
  [PresetFunc.MouseLeftButtonHold]
    Primary = "C"
    Secondary = "N"



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
		ActiveMode  []string `toml:"ActiveMode"`
		QuitMode    []string `toml:"QuitMode"`
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
				Level1      int        `toml:"Level1"`
				LevelSwitch [][]string `toml:"LevelSwitch"`
				Level2      int        `toml:"Level2"`
				Level3      int        `toml:"Level3"`
				Level4      int        `toml:"Level4"`
				Level5      int        `toml:"Level5"`
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
				LevelSwitch [][]string `toml:"LevelSwitch"`
				Level1      int        `toml:"Level1"`
				Level2      int        `toml:"Level2"`
				Level3      int        `toml:"Level3"`
				Level4      int        `toml:"Level4"`
				Level5      int        `toml:"Level5"`
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
