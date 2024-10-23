package logger

import "testing"

func Test_writeToFileWithLogLevel(t *testing.T) {
	Infof("header01", "desc01:%s", "fuck")
}
