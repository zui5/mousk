package base

import "math"

var speedLevel = 1 // the speed of you mouse movement
var mode = 0       // 0:normal, 1:control

func SetSpeedLevel(speedLevelArg int) {
	// speedLevel = speedLevelArg
	// speedLevel = int(math.Round(math.Log10(1000 * math.Pow(float64(speedLevelArg), 10))))
	// speedLevel = int(math.Log(1000 * math.Pow(float64(speedLevelArg), 5)))
	speedLevel = int(math.Pow(float64(speedLevelArg), 2))

}

func GetSpeedLevel() int {
	return speedLevel
}

func SetMode(modeArg int) {
	mode = modeArg
}
func GetMode() int {
	return mode
}
