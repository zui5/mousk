package keyboardctl

import (
	"encoding/json"
	"mousk/common/logger"
	"mousk/infra/base"
	"syscall"
	"time"
	"unsafe"
)

var (
	keyPressedStates    map[string]*KeyState
	setWindowsHookEx    = base.User32.NewProc("SetWindowsHookExW")
	getMessageW         = base.User32.NewProc("GetMessageW")
	unhookWindowsHookEx = base.User32.NewProc("UnhookWindowsHookEx")
	// todo serilize to file config
	listeningKeyReference = make(map[string]*KeyReference, 0)
)

type KeyState struct {
	Pressed            bool
	LastReleased       *time.Time
	LastPressed        *time.Time
	SecondLastReleased *time.Time
	SecondLastPressed  *time.Time
}

type KBDLLHOOKSTRUCT struct {
	VkCode   uint32
	ScanCode uint32
}
type KeyReference struct {
	Count           int
	KeyCombinations []KeyCallback
}

type KeyCallback struct {
	LastTriggerTime    time.Time
	FirstClickKeys     []string  // for example. ["CONTROL","A"]
	SecondClickKeys    []string  // for example. ["CONTROL","A"]
	Cb                 Callback2 `json:"-"`
	CbPriority         int
	withReleaseEvent   bool
	effectOnNormalMode bool
}

type Callback HookProc

// type Callback2 HookProc
type Callback2 func(wParam uintptr, vkCode, scanCode uint32) uintptr
type HookProc func(nCode int, wParam uintptr, lParam uintptr) uintptr

func registerKeyListening(cb Callback2, cbPriority int, effectOnNormal bool, withReleaseEvent bool, firstClickKeys []string, secondClickKeys []string) {
	for _, keyName := range firstClickKeys {
		if listeningKeyReference[keyName] == nil {
			listeningKeyReference[keyName] = &KeyReference{}
		}
		listeningKeyReference[keyName].Count += 1
		listeningKeyReference[keyName].KeyCombinations = append(listeningKeyReference[keyName].KeyCombinations, KeyCallback{
			FirstClickKeys:     firstClickKeys,
			SecondClickKeys:    secondClickKeys,
			Cb:                 cb,
			CbPriority:         cbPriority,
			withReleaseEvent:   withReleaseEvent,
			effectOnNormalMode: effectOnNormal,
		})
	}
	for _, keyName := range secondClickKeys {
		if listeningKeyReference[keyName] == nil {
			listeningKeyReference[keyName] = &KeyReference{}
		}
		listeningKeyReference[keyName].Count += 1
		listeningKeyReference[keyName].KeyCombinations = append(listeningKeyReference[keyName].KeyCombinations, KeyCallback{
			FirstClickKeys:     firstClickKeys,
			SecondClickKeys:    secondClickKeys,
			Cb:                 cb,
			withReleaseEvent:   withReleaseEvent,
			effectOnNormalMode: effectOnNormal,
		})
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func unRegisterKeyListening(keyNames ...string) {
	for _, vkCode := range keyNames {
		listeningKeyReference[vkCode].Count -= 1
		listeningKeyReference[vkCode].KeyCombinations = append(listeningKeyReference[vkCode].KeyCombinations, KeyCallback{
			FirstClickKeys: keyNames,
		})
	}
}

func RegisterMulti(cb Callback2, priority int, multiKeyNames ...[]string) {
	for _, keyNames := range multiKeyNames {
		registerKeyListening(cb, priority, false, false, keyNames, nil)
	}
}

func RegisterNormal(cb Callback2, priority int, keyNames ...string) {
	registerKeyListening(cb, priority, true, false, keyNames, nil)
}

func RegisterOne(cb Callback2, priority int, keyNames ...string) {
	registerKeyListening(cb, priority, false, false, keyNames, nil)
}

func RegisterDoubleClick(cb Callback2, priority int, firstClick []string, secondClick []string) {
	registerKeyListening(cb, priority, false, false, firstClick, secondClick)
}

func EffectOnNormalMode(vkCode string) bool {
	if listeningKeyReference[vkCode] != nil {
		// logger.Infof("", "effect on normal check:%+v, current:%d", listeningKeyReference[vkCode], vkCode)
		ref := listeningKeyReference[vkCode]
		for _, v := range ref.KeyCombinations {
			// logger.Infof("", "effect on normal check:%+v, %t", v.FirstClickKeys, StatusCheck(v.FirstClickKeys, 1, time.Second))
			if !v.effectOnNormalMode {
				continue
			}
			// logger.Infof("","11111:%+v, %t", v.FirstClickKeys, AllPressed(v.FirstClickKeys...))
			// if AllPressed(v.FirstClickKeys...) {
			if StatusCheckNew(v.FirstClickKeys, 1) {
				return true
			}
		}
	}
	// if AllPressed(VK_LWIN, VK_SPACE) {
	// 	return true
	// }
	// if AllPressed(VK_LALT, VK_0) {
	// 	return true
	// }
	// if AllPressed(VK_TAB, VK_SPACE) {
	// 	return true
	// }
	return false
}

func LowLevelKeyboardCallback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode == 0 {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		scanCode := kbdStruct.ScanCode
		vkCode := kbdStruct.VkCode
		keyName := GetNameByCode(vkCode)

		if wParam == WM_KEYDOWN || wParam == WM_SYSKEYDOWN {
			SetPressed(keyName)
		} else if wParam == WM_KEYUP || wParam == WM_SYSKEYUP {
			SetReleased(keyName)
			if keyName == "LSHIFT" {
				logger.Infof("", "shift dectect:%t", IsShiftPressed())
			}
		}

		// // 检查是否同时按下了 Ctrl、Shift 和 A 键
		// if Pressed(VK_LCONTROL) && Pressed(VK_LSHIFT) && Pressed(VK_A) {
		// 	logger.Infof("", "Ctrl+Shift+A keys pressed simultaneously")
		// 	os.Exit(0)
		// 	return 1
		// }

		if base.GetMode() == base.ModeNormal && !EffectOnNormalMode(keyName) {
			logger.Infof("", "%d not in control mode, mode:%d, keystatus:%s", time.Now().UnixMilli(), base.GetMode(), keyName)
			return 0
		}

		if listeningKeyReference[keyName] != nil {
			ref := listeningKeyReference[keyName]
			// TODO : ref.KeyCombinations中满足多个快捷键组合时，仅优先执行其中按键数量最多的
			// k | ctrl+k, 优先执行ctrl+k

			satisfiedCallback := make([]KeyCallback, 0)
			for _, v := range ref.KeyCombinations {
				// if AllPressed(v.FirstClickKeys...) {
				if StatusCheckNew(v.FirstClickKeys, 1) {
					satisfiedCallback = append(satisfiedCallback, v)
				}
				if v.withReleaseEvent && StatusCheckNew(v.FirstClickKeys, 0) {
					satisfiedCallback = append(satisfiedCallback, v)
				}
			}

			// return 1 is importantA
			// 如果没有匹配的快捷键，返回1
			if len(satisfiedCallback) == 0 {
				return 0
			}

			// order by priority
			mostKeyNumCallback := satisfiedCallback[0]

			for _, v := range satisfiedCallback {
				logger.Infof("", "all keycallback:%+v", v.FirstClickKeys)
				if len(v.FirstClickKeys) > len(mostKeyNumCallback.FirstClickKeys) {
					mostKeyNumCallback = v
				} else if len(v.FirstClickKeys) == len(mostKeyNumCallback.FirstClickKeys) {
					if v.CbPriority > mostKeyNumCallback.CbPriority {
						mostKeyNumCallback = v
					}

				} else {
				}
			}
			logger.Infof("", "most keycallback:%+v", mostKeyNumCallback.FirstClickKeys)
			mostKeyNumCallback.Cb(wParam, vkCode, scanCode)
			return 1
		}
		return 0
	}
	return 0
}

func RawKeyboardListener(cb Callback) {

	keyPressedStates = make(map[string]*KeyState)

	hookProc := HookProc(cb)

	hookID, _, _ := setWindowsHookEx.Call(
		uintptr(WH_KEYBOARD_LL),
		syscall.NewCallback(hookProc),
		0,
		0,
	)

	if hookID == 0 {
		logger.Infof("", "Failed to set hook")
		return
	}

	logger.Infof("", "Hook set, waiting for events...")

	defer unhookWindowsHookEx.Call(hookID)

	var msg uintptr
	for {
		_, _, _ = getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}

}

func Pressed(vkCode string) bool {
	// logger.Infof("",vkCode, keyPressedStates[vkCode])
	if keyPressedStates[vkCode] == nil {
		keyPressedStates[vkCode] = nilKeyState()
	}
	return keyPressedStates[vkCode].Pressed
}

// deprecated
func AllPressed(vkCodes ...string) bool {
	if vkCodes == nil {
		return true
	}
	allPressed := true
	for _, v := range vkCodes {
		logger.Infof("", "keys %d , %t", v, Pressed(v))
		allPressed = allPressed && Pressed(v)
	}
	return allPressed
}

func SetPressed(vkCode string) {
	if keyPressedStates[vkCode] == nil {
		keyPressedStates[vkCode] = nilKeyState()
	}
	keyPressedStates[vkCode].Pressed = true

	currTime := time.Now()
	keyPressedStates[vkCode].LastPressed = &currTime
	keyPressedStates[vkCode].SecondLastPressed = keyPressedStates[vkCode].LastPressed
	logger.Infof("", "Key pressed (VK code): %d, last pressed: %s", vkCode, currTime)
}

func nilKeyState() *KeyState {
	return &KeyState{
		Pressed:            false,
		LastReleased:       nil,
		LastPressed:        nil,
		SecondLastPressed:  nil,
		SecondLastReleased: nil,
	}
}

func SetReleased(vkCode string) {
	if keyPressedStates[vkCode] == nil {
		keyPressedStates[vkCode] = nilKeyState()
	}
	keyPressedStates[vkCode].Pressed = false
	currTime := time.Now()
	keyPressedStates[vkCode].SecondLastReleased = keyPressedStates[vkCode].LastReleased
	keyPressedStates[vkCode].LastReleased = &currTime
	logger.Infof("", "Key released (VK code): %d, last released: %s", vkCode, currTime)
}

func RegisterWithReleaseEventMulti(cb Callback2, prority int, mulitiVkCodes ...[]string) {
	// switch keyAction {
	// case WM_KEYDOWN:

	// 	break
	// default:
	// 	logger.Infof("","method not develop yet")
	// 	return
	// }
	for _, vkCodes := range mulitiVkCodes {
		registerKeyListening(cb, prority, false, true, vkCodes, nil)
	}
}

func StatusCheckNew(keyNames []string, status int) bool {
	if len(keyNames) == 0 {
		return false
	}
	for _, name := range keyNames {
		if status == 1 {
			if !Pressed(name) {
				return false
			}
		} else {
			if Pressed(name) {
				return false
			}
		}
	}
	return true
}

func PrintAllKeys() string {
	keysRaw, _ := json.Marshal(listeningKeyReference)
	logger.Infof("", "all key refrence", string(keysRaw))
	return string(keysRaw)
}
