package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	_ "embed"

	"github.com/BurntSushi/toml"
)

var settingsVar Settings

type Mode int

var (
	ModeLoadFromDefault Mode = Mode(0)
	ModeLoadFromUser    Mode = Mode(1)
)

//go:embed conf/default.toml
var defaultConfigFile []byte

var defaultFilePath = "./conf/default.toml"

// var filePath = "./confs.toml"

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

// type Settings struct {
// 	Mouse struct {
// 		Fast MouseConfig
// 		Slow MouseConfig
// 	}
// }

type Settings struct {
	ENV             string  `toml:"ENV"`
	VER             float64 `toml:"VER"`
	StartOnSystemUp bool    `toml:"StartOnSystemUp"`
	PresetFunc      struct {
		HelpPane          []string `toml:"HelpPane"`
		ForceQuit         []string `toml:"ForceQuit"`
		ActiveMode        []string `toml:"ActiveMode"`
		ResetSetting      []string `toml:"ResetSetting"`
		ToggleControlMode []string `toml:"ToggleControlMode"`
		TmpQuitMode       []string `toml:"TmpQuitMode"`
		QuitMode          []string `toml:"QuitMode"`
		OpenSetting       []string `toml:"OpenSetting"`
		MouseMove         struct {
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
				LevelSwitch []int `toml:"LevelSwitch"`
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
				LevelSwitch []int `toml:"LevelSwitch"`
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

func init() {
	if initConfigFile() != nil {
		log.Fatalf("init config file error")
		return
	}
}

func LoadSettingsFromFile(mode Mode) error {
	var filePath string
	var err error
	if mode == 1 {
		_, filePath, err = getUserConfigPath()
		fmt.Println("config file path:" + filePath)
		if _, err = toml.DecodeFile(filePath, &settingsVar); err != nil {
			log.Fatalf("decode config file error:%+v", err)
			return fmt.Errorf("decoding TOML:%+v", err)
		}
	} else if mode == 0 {
		// toml.DecodeReader(io.ByteReader(), v any)
		if _, err = toml.NewDecoder(bytes.NewReader(defaultConfigFile)).Decode(&settingsVar); err != nil {
			log.Fatalf("decode default config file error:%+v", err)
			return fmt.Errorf("decoding TOML:%+v", err)
		}
	} else {
		return fmt.Errorf("not valid mode:%d", mode)
	}

	// 输出生成的结构体
	// log.Infof("", "%#v", settingsVar)
	log.Printf("setting var:%#v", settingsVar)
	return nil
}

func GetSettings() *Settings {
	return &settingsVar
}

func WriteSettings() error {
	_, filePath, err := getUserConfigPath()
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

// func getEmbededConfigPath() string {
// 	return defaultFilePath
// }

func getUserConfigPath() (string, string, error) {
	// 获取当前用户的主目录
	usr, err := user.Current()
	if err != nil {
		return "", "", err
	}

	// 构建文件路径
	configDir := filepath.Join(usr.HomeDir, ".config", "mousk")
	filePath := filepath.Join(configDir, "confs.toml")
	return configDir, filePath, nil
}

func initConfigFile() error {
	// 构建文件路径
	configDir, filePath, err := getUserConfigPath()
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
		// logger.Infof("", "Config file created at:", filePath)
		log.Printf("Config file created at:%s", filePath)
	} else {
		err = LoadSettingsFromFile(ModeLoadFromUser)
		if err != nil {
			return err
		}
		// logger.Infof("", "Config file already exists at:", filePath)
		log.Printf("Config file already exists at:%s", filePath)
	}
	return nil
}

func RestoreSettings() error {
	if err := LoadSettingsFromFile(ModeLoadFromDefault); err != nil {
		return fmt.Errorf("decoding TOML:%+v", err)
	}
	// 输出生成的结构体
	// logger.Infof("", "restore setting:%#v", settingsVar)
	log.Printf("restore setting:%#v", settingsVar)
	if err := WriteSettings(); err != nil {
		return err
	}
	return nil
}
