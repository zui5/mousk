package keyboardctl

var validKeycodeNameMap = map[uint32]string{
	// common use keycode
	9:   "TAB",     //Tab
	16:  "SHIFT",   //Shift
	17:  "CONTROL", //Ctrl
	20:  "CAPITAL", //Caps Lock
	27:  "ESCAPE",  //Esc
	32:  "SPACE",   //Space
	48:  "0",
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "Q",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	160: "LSHIFT",   //左SHIFT键
	161: "RSHIFT",   //右SHIFT键
	162: "LCONTROL", //左CONTROL键
	163: "RCONTROL", //右CONTROL键
	188: "COMMA",

	// not really in use
	// 1:  "LBUTTON",  //鼠标左键
	// 2:  "RBUTTON",  //鼠标右键
	// 3:  "CANCEL",   //Cancel
	// 4:  "MBUTTON",  //鼠标中键
	// 5:  "XBUTTON1", //X1鼠标按钮
	// 6:  "XBUTTON2", //X2鼠标按钮
	// 8:  "BACK",     //Backspace
	// 12: "CLEAR",    //Clear
	// 13: "RETURN",   //Enter
	// 18: "ALT", //ALt
	// 19: "PAUSE",    //Pause
	// 21:  "KANA",       //IME假名模式
	// 21:  "HANGUL",     //IME Hanguel模式（保持兼容性；使用"HANGUL",）
	// 21:  "HANGEUL",    //IME韩文模式
	// 23:  "JUNJA",      //IME Junja模式
	// 24:  "FINAL",      //IME最终模式
	// 25:  "HANJA",      //IME Hanja模式
	// 25:  "HANJI",      //IME汉字模式
	// 28:  "CONVERT",    //输入法转换
	// 29:  "NONCONVERT", //IME不可转换
	// 30:  "ACCETP",     //输入法接受
	// 31:  "MODECHANGE", //输入法模式更改请求
	// 33:  "PRIOR",      //Page Up
	// 34:  "NEXT",       //Page Down
	// 35:  "END",        //End
	// 36:  "HOME",       //Home
	// 37:  "LEFT",       //Left Arrow
	// 38:  "UP",         //Up Arrow
	// 39:  "RIGHT",      //Right Arrow
	// 40:  "DOWN",       //Down Arrow
	// 41:  "SELECT",     //Select
	// 42:  "PRINT",      //Print
	// 43:  "EXECUTE",    //Execute
	// 44:  "SNAPSHOT",   //Snapshot
	// 45:  "INSERT",     //Insert
	// 46:  "DELETE",     //Delete
	// 47:  "HELP",       //Help
	// 91:  "LWIN",      //左Windows键（自然键盘）
	// 92:  "RWIN",      //右Windows键（自然键盘）
	// 93:  "APPS",      //应用程序键（自然键盘）
	// 95:  "SLEEP",     //电脑睡眠键
	// 96:  "NUMPAD0",   //小键盘0
	// 97:  "NUMPAD1",   //小键盘1
	// 98:  "NUMPAD2",   //小键盘2
	// 99:  "NUMPAD3",   //小键盘3
	// 100: "NUMPAD4",   //小键盘4
	// 101: "NUMPAD5",   //小键盘5
	// 102: "NUMPAD6",   //小键盘6
	// 103: "NUMPAD7",   //小键盘7
	// 104: "NUMPAD8",   //小键盘8
	// 105: "NUMPAD9",   //小键盘9
	// 106: "MULTIPLY",  //小键盘 *
	// 107: "ADD",       //小键盘 +
	// 108: "SEPARATOR", //小键盘 Enter
	// 109: "SUBTRACT",  //小键盘 -
	// 110: "DECIMAL",   //小键盘 .
	// 111: "DIVIDE",    //小键盘 /
	// 112: "F1",        //F1
	// 113: "F2",        //F2
	// 114: "F3",        //F3
	// 115: "F4",        //F4
	// 116: "F5",        //F5
	// 117: "F6",        //F6
	// 118: "F7",        //F7
	// 119: "F8",        //F8
	// 120: "F9",        //F9
	// 121: "F10",       //F10
	// 122: "F11",       //F11
	// 123: "F12",       //F12
	// 124: "F13",
	// 125: "F14",
	// 126: "F15",
	// 127: "F16",
	// 128: "F17",
	// 129: "F18",
	// 130: "F19",
	// 131: "F20",
	// 132: "F21",
	// 133: "F22",
	// 134: "F23",
	// 135: "F24",
	// 144: "MUMLOCK",             //Num Lock
	// 145: "SCROLL",              //Scroll
	164: "LALT", //左ALT键
	165: "RALT", //右ALT键
	// 166: "BROWSER_BACK",        //浏览器后退键
	// 167: "BROWSER_FORWARD",     //浏览器前进键
	// 168: "BROWSER_REFRESH",     //浏览器刷新键
	// 169: "BROWSER_STOP",        //浏览器停止键
	// 170: "BROWSER_SEARCH",      //浏览器搜索键
	// 171: "BROWSER_FAVORITES",   //浏览器收藏夹键
	// 172: "BROWSER_HOME",        //浏览器开始和主页键
	// 173: "VOLUME_MUTE",         //音量静音键
	// 174: "VOLUME_DOWN",         //降低音量键
	// 175: "VOLUME_UP",           //调高音量键
	// 176: "MEDIA_NEXT_TRACK",    //下一曲目键
	// 177: "MEDIA_PREV_TRACK",    //上一曲目键
	// 178: "MEDIA_STOP",          //停止媒体键
	// 179: "MEDIA_PLAY_PAUSE",    //播放/暂停媒体键
	// 180: "LAUNCH_MAIL",         //启动邮件密钥
	// 181: "LAUNCH_MEDIA_SELECT", //选择媒体密钥
	// 182: "LAUNCH_APP1",         //启动应用程序1键
	// 183: "LAUNCH_APP2",         //启动应用程序2键
	// 186: "OEM_1",               //;:
	// 187: "OEM_PLUS",            //:+
	// 189: "OEM_MINUS", //-_
	190: "PERIOD",
	// 191: "OEM_2", ///?
	// 192: "OEM_3", //`~
	// 219: "OEM_4", //[{
	// 220: "OEM_5", //\ |
	// 221: "OEM_6", //]}
	// 222: "OEM_7", //'"
	// 223: "OEM_8",
	// 226: "OEM_102",
	// 227: "ICO_HELP",
	// 228: "ICO_00",
	// 229: "PROCESSKEY",
	// 230: "ICO_CLEAR",
	// 231: "PACKET",
	// 233: "OEM_RESET",
	// 234: "OEM_JUMP",
	// 235: "OEM_PA1",
	// 236: "OEM_PA2",
	// 237: "OEM_PA3",
	// 238: "OEM_WSCTRL",
	// 239: "OEM_CUSEL",
	// 240: "OEM_ATTN",
	// 241: "OEM_FINISH",
	// 242: "OEM_COPY",
	// 243: "OEM_AUTO",
	// 244: "OEM_ENLW",
	// 245: "OEM_BACKTAB",
	// 246: "ATTN",
	// 247: "CRSEL",
	// 248: "EXSEL",
	// 249: "EREOF",
	// 250: "PLAY",
	// 251: "ZOOM",
	// 252: "NONAME",
	// 253: "PA1",
	// 254: "OEM_CLEAR",
}

var validNameKeycodeMap = map[string]uint32{
	"LBUTTON":             1,  //鼠标左键
	"RBUTTON":             2,  //鼠标右键
	"CANCEL":              3,  //Cancel
	"MBUTTON":             4,  //鼠标中键
	"XBUTTON1":            5,  //X1鼠标按钮
	"XBUTTON2":            6,  //X2鼠标按钮
	"BACK":                8,  //Backspace
	"TAB":                 9,  //Tab
	"CLEAR":               12, //Clear
	"RETURN":              13, //Enter
	"SHIFT":               16, //Shift
	"CONTROL":             17, //Ctrl
	"ALT":                 18, //ALt
	"PAUSE":               19, //Pause
	"CAPITAL":             20, //Caps Lock
	"KANA":                21, //IME假名模式
	"HANGUL":              21, //IME Hanguel模式（保持兼容性；使用"HANGUL",）
	"HANGEUL":             21, //IME韩文模式
	"JUNJA":               23, //IME Junja模式
	"FINAL":               24, //IME最终模式
	"HANJA":               25, //IME Hanja模式
	"HANJI":               25, //IME汉字模式
	"ESCAPE":              27, //Esc
	"CONVERT":             28, //输入法转换
	"NONCONVERT":          29, //IME不可转换
	"ACCETP":              30, //输入法接受
	"MODECHANGE":          31, //输入法模式更改请求
	"SPACE":               32, //Space
	"PRIOR":               33, //Page Up
	"NEXT":                34, //Page Down
	"END":                 35, //End
	"HOME":                36, //Home
	"LEFT":                37, //Left Arrow
	"UP":                  38, //Up Arrow
	"RIGHT":               39, //Right Arrow
	"DOWN":                40, //Down Arrow
	"SELECT":              41, //Select
	"PRINT":               42, //Print
	"EXECUTE":             43, //Execute
	"SNAPSHOT":            44, //Snapshot
	"INSERT":              45, //Insert
	"DELETE":              46, //Delete
	"HELP":                47, //Help
	"0":                   48,
	"1":                   49,
	"2":                   50,
	"3":                   51,
	"4":                   52,
	"5":                   53,
	"6":                   54,
	"7":                   55,
	"8":                   56,
	"9":                   57,
	"A":                   65,
	"B":                   66,
	"C":                   67,
	"D":                   68,
	"E":                   69,
	"F":                   70,
	"G":                   71,
	"H":                   72,
	"I":                   73,
	"J":                   74,
	"K":                   75,
	"L":                   76,
	"M":                   77,
	"N":                   78,
	"O":                   79,
	"P":                   80,
	"Q":                   81,
	"R":                   82,
	"S":                   83,
	"T":                   84,
	"U":                   85,
	"V":                   86,
	"W":                   87,
	"X":                   88,
	"Y":                   89,
	"Z":                   90,
	"LWIN":                91,  //左Windows键（自然键盘）
	"RWIN":                92,  //右Windows键（自然键盘）
	"APPS":                93,  //应用程序键（自然键盘）
	"SLEEP":               95,  //电脑睡眠键
	"NUMPAD0":             96,  //小键盘0
	"NUMPAD1":             97,  //小键盘1
	"NUMPAD2":             98,  //小键盘2
	"NUMPAD3":             99,  //小键盘3
	"NUMPAD4":             100, //小键盘4
	"NUMPAD5":             101, //小键盘5
	"NUMPAD6":             102, //小键盘6
	"NUMPAD7":             103, //小键盘7
	"NUMPAD8":             104, //小键盘8
	"NUMPAD9":             105, //小键盘9
	"MULTIPLY":            106, //小键盘 *
	"ADD":                 107, //小键盘 +
	"SEPARATOR":           108, //小键盘 Enter
	"SUBTRACT":            109, //小键盘 -
	"DECIMAL":             110, //小键盘 .
	"DIVIDE":              111, //小键盘 /
	"F1":                  112, //F1
	"F2":                  113, //F2
	"F3":                  114, //F3
	"F4":                  115, //F4
	"F5":                  116, //F5
	"F6":                  117, //F6
	"F7":                  118, //F7
	"F8":                  119, //F8
	"F9":                  120, //F9
	"F10":                 121, //F10
	"F11":                 122, //F11
	"F12":                 123, //F12
	"F13":                 124,
	"F14":                 125,
	"F15":                 126,
	"F16":                 127,
	"F17":                 128,
	"F18":                 129,
	"F19":                 130,
	"F20":                 131,
	"F21":                 132,
	"F22":                 133,
	"F23":                 134,
	"F24":                 135,
	"MUMLOCK":             144, //Num Lock
	"SCROLL":              145, //Scroll
	"LSHIFT":              160, //左SHIFT键
	"RSHIFT":              161, //右SHIFT键
	"LCONTROL":            162, //左CONTROL键
	"RCONTROL":            163, //右CONTROL键
	"LALT":                164, //左ALT键
	"RALT":                165, //右ALT键
	"BROWSER_BACK":        166, //浏览器后退键
	"BROWSER_FORWARD":     167, //浏览器前进键
	"BROWSER_REFRESH":     168, //浏览器刷新键
	"BROWSER_STOP":        169, //浏览器停止键
	"BROWSER_SEARCH":      170, //浏览器搜索键
	"BROWSER_FAVORITES":   171, //浏览器收藏夹键
	"BROWSER_HOME":        172, //浏览器开始和主页键
	"VOLUME_MUTE":         173, //音量静音键
	"VOLUME_DOWN":         174, //降低音量键
	"VOLUME_UP":           175, //调高音量键
	"MEDIA_NEXT_TRACK":    176, //下一曲目键
	"MEDIA_PREV_TRACK":    177, //上一曲目键
	"MEDIA_STOP":          178, //停止媒体键
	"MEDIA_PLAY_PAUSE":    179, //播放/暂停媒体键
	"LAUNCH_MAIL":         180, //启动邮件密钥
	"LAUNCH_MEDIA_SELECT": 181, //选择媒体密钥
	"LAUNCH_APP1":         182, //启动应用程序1键
	"LAUNCH_APP2":         183, //启动应用程序2键
	"OEM_1":               186, //;:
	"OEM_PLUS":            187, //:+
	// "OEM_COMMA":           188,
	"COMMA":       188,
	"OEM_MINUS":   189, //-_
	"PERIOD":      190,
	"OEM_2":       191, ///?
	"OEM_3":       192, //`~
	"OEM_4":       219, //[{
	"OEM_5":       220, //\ |
	"OEM_6":       221, //]}
	"OEM_7":       222, //'"
	"OEM_8":       223,
	"OEM_102":     226,
	"ICO_HELP":    227,
	"ICO_00":      228,
	"PROCESSKEY":  229,
	"ICO_CLEAR":   230,
	"PACKET":      231,
	"OEM_RESET":   233,
	"OEM_JUMP":    234,
	"OEM_PA1":     235,
	"OEM_PA2":     236,
	"OEM_PA3":     237,
	"OEM_WSCTRL":  238,
	"OEM_CUSEL":   239,
	"OEM_ATTN":    240,
	"OEM_FINISH":  241,
	"OEM_COPY":    242,
	"OEM_AUTO":    243,
	"OEM_ENLW":    244,
	"OEM_BACKTAB": 245,
	"ATTN":        246,
	"CRSEL":       247,
	"EXSEL":       248,
	"EREOF":       249,
	"PLAY":        250,
	"ZOOM":        251,
	"NONAME":      252,
	"PA1":         253,
	"OEM_CLEAR":   254,
}

func GetCodeByName(name string) uint32 {
	if val, ok := validNameKeycodeMap[name]; ok {
		return val
	}
	return 0
}

func ExportAllCodes() [][]uint32 {
	codes := make([][]uint32, 0)
	for _, v := range validNameKeycodeMap {
		codes = append(codes, []uint32{v})
	}
	return codes
}

func GetCodesByNames(names []string) []uint32 {
	codes := make([]uint32, 0)
	for _, v := range names {
		code := GetCodeByName(v)
		if code == 0 {
			return nil
		}
		codes = append(codes, code)
	}
	return codes
}

func GetNameByCode(code uint32) string {
	if val, ok := validKeycodeNameMap[code]; ok {
		return val
	}
	return ""
}

func GetNamesByCodes(codes []uint32) []string {
	names := make([]string, 0)
	for _, v := range codes {
		name := GetNameByCode(v)
		if name == "" {
			return nil
		}
		names = append(names, name)
	}
	return names
}
