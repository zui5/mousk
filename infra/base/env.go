package base

import "math"

var moveSpeedLevel = 1   // the speed of you mouse movement
var scrollSpeedLevel = 1 // the speed of you mouse scroll
var mode = 0             // 0:normal, 1:control

func SetMoveSpeedLevel(speedLevelArg int) {
	// speedLevel = speedLevelArg
	// speedLevel = int(math.Round(math.Log10(1000 * math.Pow(float64(speedLevelArg), 10))))
	// speedLevel = int(math.Log(1000 * math.Pow(float64(speedLevelArg), 5)))
	moveSpeedLevel = int(math.Pow(float64(speedLevelArg), 2))
}

func SetScrollSpeedLevel(speedLevelArg int) {
	// speedLevel = speedLevelArg
	// speedLevel = int(math.Round(math.Log10(1000 * math.Pow(float64(speedLevelArg), 10))))
	// speedLevel = int(math.Log(1000 * math.Pow(float64(speedLevelArg), 5)))
	scrollSpeedLevel = int(math.Pow(float64(speedLevelArg), 2))
}

func GetMoveSpeedLevel() int {
	return moveSpeedLevel
}

func GetScrollSpeedLevel() int {
	return scrollSpeedLevel
}

func SetMode(modeArg int) {
	mode = modeArg
}
func GetMode() int {
	return mode
}
