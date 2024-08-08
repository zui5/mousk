package keyboardctl

import (
	"encoding/json"
	"fmt"
	"mousek/infra/base"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var (
	keyPressedStates    map[uint32]*KeyState
	setWindowsHookEx    = base.User32.NewProc("SetWindowsHookExW")
	getMessageW         = base.User32.NewProc("GetMessageW")
	unhookWindowsHookEx = base.User32.NewProc("UnhookWindowsHookEx")
	// todo serilize to file config
	listeningKeyReference = make(map[uint32]*KeyReference, 0)
)

type KeyState struct {
	Pressed            bool
	LastReleased       *time.Time
	SecondLastReleased *time.Time
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
	FirstClickKeys     []uint32  // for example. []{"ctrl","a"}
	SecondClickKeys    []uint32  // for example. []{"ctrl","a"}
	Cb                 Callback2 `json:"-"`
	withReleaseEvent   bool
	effectOnNormalMode bool
}

type Callback HookProc

// type Callback2 HookProc
type Callback2 func(wParam uintptr, vkCode, scanCode uint32) uintptr
type HookProc func(nCode int, wParam uintptr, lParam uintptr) uintptr

func registerKeyListening(cb Callback2, effectOnNormal bool, withReleaseEvent bool, firstClickVkCodes []uint32, secondClickVkCodes []uint32) {
	for _, v := range firstClickVkCodes {
		if listeningKeyReference[v] == nil {
			listeningKeyReference[v] = &KeyReference{}
		}
		listeningKeyReference[v].Count += 1
		listeningKeyReference[v].KeyCombinations = append(listeningKeyReference[v].KeyCombinations, KeyCallback{
			FirstClickKeys:     firstClickVkCodes,
			SecondClickKeys:    secondClickVkCodes,
			Cb:                 cb,
			withReleaseEvent:   withReleaseEvent,
			effectOnNormalMode: effectOnNormal,
		})
	}
	for _, v := range secondClickVkCodes {
		if listeningKeyReference[v] == nil {
			listeningKeyReference[v] = &KeyReference{}
		}
		listeningKeyReference[v].Count += 1
		listeningKeyReference[v].KeyCombinations = append(listeningKeyReference[v].KeyCombinations, KeyCallback{
			FirstClickKeys:     firstClickVkCodes,
			SecondClickKeys:    secondClickVkCodes,
			Cb:                 cb,
			withReleaseEvent:   withReleaseEvent,
			effectOnNormalMode: effectOnNormal,
		})
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func unRegisterKeyListening(vkCodes ...uint32) {
	for _, v := range vkCodes {
		listeningKeyReference[v].Count -= 1
		listeningKeyReference[v].KeyCombinations = append(listeningKeyReference[v].KeyCombinations, KeyCallback{
			FirstClickKeys: vkCodes,
		})
	}

}

func RegisterMulti(cb Callback2, mulitiVkCodes ...[]uint32) {
	// switch keyAction {
	// case WM_KEYDOWN:

	// 	break
	// default:
	// 	fmt.Println("method not develop yet")
	// 	return
	// }
	for _, vkCodes := range mulitiVkCodes {
		registerKeyListening(cb, false, false, vkCodes, nil)
	}
}

func RegisterNormal(cb Callback2, vkCodes ...uint32) {
	registerKeyListening(cb, true, false, vkCodes, nil)
}

func RegisterOne(cb Callback2, vkCodes ...uint32) {
	registerKeyListening(cb, false, false, vkCodes, nil)
}

func RegisterDoubleClick(cb Callback2, firstClick []uint32, secondClick []uint32) {
	registerKeyListening(cb, false, false, firstClick, secondClick)
}

func EffectOnNormalMode(vkCode uint32) bool {
	if listeningKeyReference[vkCode] != nil {
		ref := listeningKeyReference[vkCode]
		for _, v := range ref.KeyCombinations {
			if !v.effectOnNormalMode {
				continue
			}
			fmt.Printf("11111:%+v, %t", v.FirstClickKeys, AllPressed(v.FirstClickKeys...))
			if AllPressed(v.FirstClickKeys...) {
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

		if wParam == WM_KEYDOWN || wParam == WM_SYSKEYDOWN {
			// TODO
			SetPressed(vkCode)
			fmt.Printf("Key pressed (VK code): %d, Scan code: %d\n", vkCode, scanCode)
		} else if wParam == WM_KEYUP || wParam == WM_SYSKEYUP {
			SetReleased(vkCode)
			fmt.Printf("Key released (VK code): %d, Scan code: %d\n", vkCode, scanCode)
		}

		// 检查是否同时按下了 Ctrl、Shift 和 A 键
		if Pressed(VK_LCONTROL) && Pressed(VK_LSHIFT) && Pressed(VK_A) {
			fmt.Println("Ctrl+Shift+A keys pressed simultaneously")
			os.Exit(0)
			return 1
		}

		if base.GetMode() == base.ModeNormal && !EffectOnNormalMode(vkCode) {
			fmt.Printf("%d not in control mode, mode:%d, keystatus:%d\n", time.Now().UnixMilli(), base.GetMode(), wParam)
			return 0
		}

		if listeningKeyReference[vkCode] != nil {
			ref := listeningKeyReference[vkCode]
			// TODO : ref.KeyCombinations中满足多个快捷键组合时，仅优先执行其中按键数量最多的
			// k | ctrl+k, 优先执行ctrl+k

			satisfiedCallback := make([]KeyCallback, 0)
			for _, v := range ref.KeyCombinations {
				if AllPressed(v.FirstClickKeys...) {
					satisfiedCallback = append(satisfiedCallback, v)
				}
				if v.withReleaseEvent && AllReleased(time.Second, v.FirstClickKeys...) {
					satisfiedCallback = append(satisfiedCallback, v)
				}
			}

			// return 1 is importantA
			// 如果没有匹配的快捷键，返回1
			if len(satisfiedCallback) == 0 {
				return 0
			}

			mostKeyNumCallback := satisfiedCallback[0]
			for _, v := range satisfiedCallback {
				if len(v.FirstClickKeys) >= len(mostKeyNumCallback.FirstClickKeys) {
					mostKeyNumCallback = v
				}
			}
			return mostKeyNumCallback.Cb(wParam, vkCode, scanCode)

			// for _, v := range ref.KeyCombinations {
			// 	if AllPressed(v.Keys...) {
			// 		v.Cb(wParam, vkCode, scanCode)
			// 	}
			// 	if v.withReleaseEvent && AllReleased(time.Second, v.Keys...) {
			// 		v.Cb(wParam, vkCode, scanCode)
			// 	}
			// }
		}
		return 0

		// // 如果按下了 'Q' 键，退出程序
		// if Pressed(VK_Q) {
		// 	os.Exit(0)
		// }

		// 在这里添加你的其他逻辑

		// return CallNextHookEx(0, nCode, wParam, lParam)
	}
	// return CallNextHookEx(0, nCode, wParam, lParam)
	return 0
}

func KeyboardCallback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vkCode := kbdStruct.VkCode

		if wParam == WM_KEYDOWN {
			SetPressed(vkCode)
			fmt.Printf("Key pressed (VK code): %d\n", vkCode)
		} else if wParam == WM_KEYUP {
			SetReleased(vkCode)
			fmt.Printf("Key released (VK code): %d\n", vkCode)
		}

		// 检查是否同时按下了 Ctrl、Shift 和 A 键
		if Pressed(VK_LCONTROL) && Pressed(VK_LSHIFT) && Pressed(VK_A) {
			fmt.Println("Ctrl+Shift+A keys pressed simultaneously")
		}

		// 如果按下了 'Q' 键，退出程序
		if Pressed(VK_Q) {
			os.Exit(0)
		}
		return 1
	}
	return 0
}

func RawKeyboardListener(cb Callback) {

	keyPressedStates = make(map[uint32]*KeyState)

	hookProc := HookProc(cb)

	hookID, _, _ := setWindowsHookEx.Call(
		uintptr(WH_KEYBOARD_LL),
		syscall.NewCallback(hookProc),
		0,
		0,
	)

	if hookID == 0 {
		fmt.Println("Failed to set hook")
		return
	}

	fmt.Println("Hook set, waiting for events...")

	defer unhookWindowsHookEx.Call(hookID)

	var msg uintptr
	for {
		_, _, _ = getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}

}

func Pressed(vkCode uint32) bool {
	// fmt.Println(vkCode, keyPressedStates[vkCode])
	if keyPressedStates[vkCode] == nil {
		keyPressedStates[vkCode] = &KeyState{false, nil, nil}
	}
	return keyPressedStates[vkCode].Pressed
}

// TODO 要设置两个按键的点击间隔过滤
func AllPressed(vkCodes ...uint32) bool {
	if vkCodes == nil {
		return true
	}
	allPressed := true
	for _, v := range vkCodes {
		fmt.Printf("keys %d , %t\n", v, Pressed(v))
		allPressed = allPressed && Pressed(v)
	}
	return allPressed
}

func SetPressed(vkCode uint32) {
	if keyPressedStates[vkCode] == nil {
		keyPressedStates[vkCode] = &KeyState{false, nil, nil}
	}
	keyPressedStates[vkCode].Pressed = true
}

func SetReleased(vkCode uint32) {
	if keyPressedStates[vkCode] == nil {
		keyPressedStates[vkCode] = &KeyState{false, nil, nil}
	}
	keyPressedStates[vkCode].Pressed = false
	currTime := time.Now()
	keyPressedStates[vkCode].SecondLastReleased = keyPressedStates[vkCode].LastReleased
	keyPressedStates[vkCode].LastReleased = &currTime
}

func RegisterWithReleaseEventMulti(cb Callback2, mulitiVkCodes ...[]uint32) {
	// switch keyAction {
	// case WM_KEYDOWN:

	// 	break
	// default:
	// 	fmt.Println("method not develop yet")
	// 	return
	// }
	for _, vkCodes := range mulitiVkCodes {
		registerKeyListening(cb, false, true, vkCodes, nil)
	}
}

func AllReleased(durationBetween time.Duration, vkCodes ...uint32) bool {
	if vkCodes == nil {
		return true
	}

	var maxLastReleasedTime *time.Time = nil
	var minLastReleasedTime *time.Time = nil
	for _, v := range vkCodes {
		keyState := keyPressedStates[v]
		if keyState.Pressed || keyState.LastReleased == nil {
			return false
		}

		if maxLastReleasedTime == nil {
			maxLastReleasedTime = keyState.LastReleased
		} else {
			if maxLastReleasedTime.Sub(*keyState.LastReleased) < 0 {
				maxLastReleasedTime = keyState.LastReleased
			}
		}
		if minLastReleasedTime == nil {
			minLastReleasedTime = keyState.LastReleased
		} else {
			if minLastReleasedTime.Sub(*keyState.LastReleased) > 0 {
				minLastReleasedTime = keyState.LastReleased
			}
		}
	}
	durationVal := minLastReleasedTime.Sub(*maxLastReleasedTime)
	if durationVal < 0 {
		durationVal = -1 * durationVal
	}
	if durationVal > durationBetween {
		return false
	}
	return true
}

func PrintAllKeys() string {
	keysRaw, _ := json.Marshal(listeningKeyReference)
	return string(keysRaw)
}
