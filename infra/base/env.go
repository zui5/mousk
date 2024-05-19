package base

var speedLevel = 1 // the speed of you mouse movement
var mode = 0       // 0:normal, 1:control

func SetSpeedLevel(speedLevelArg int) {
	speedLevel = speedLevelArg
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
