package logger

import (
	"fmt"
	"mousk/infra/base"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	fileName = "./dist/mousk.log"
)

var (
	instance     *zap.SugaredLogger = nil
	consoleOuput                    = false
)

func SetConsoleOutput(b bool) {
	consoleOuput = b
}

func init() {
	consoleOuput = !base.IsProduct()

	lumberjacklogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	defer lumberjacklogger.Close()

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// fileEncoder := zapcore.NewJSONEncoder(config)
	fileEncoder := zapcore.NewConsoleEncoder(config)

	// logFile, _ := os.OpenFile("./log-debug-zap.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //日志记录debug信息
	// errFile, _ := os.OpenFile("./log-err-zap.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //日志记录error信息
	fileCore := zapcore.AddSync(lumberjacklogger)
	// multiSyncer := zapcore.NewMultiWriteSyncer(fileCore, consoleCore)

	coreList := make([]zapcore.Core, 0)
	coreList = append(coreList, zapcore.NewCore(fileEncoder, fileCore, zap.DebugLevel))
	if consoleOuput {
		consoleCore := zapcore.AddSync(os.Stdout)
		coreList = append(coreList, zapcore.NewCore(fileEncoder, consoleCore, zap.DebugLevel))
	}
	teecore := zapcore.NewTee(coreList...)
	instance = zap.New(teecore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	defer instance.Sync()

	// // 测试分割日志
	// for i := 0; i < 118000; i++ {
	// 	logger.With(
	// 		zap.String("url", fmt.Sprintf("www.test%d.com", i)),
	// 		zap.String("name", "jimmmyr"),
	// 		zap.Int("age", 23),
	// 		zap.String("agradege", "no111-000222"),
	// 	).Info("test info ")
	// }
}

func Infof(header, template string, args ...interface{}) {
	instance.Infof(fmt.Sprintf("%s %s", header, template), args...)
}

func Warnf(header, template string, args ...interface{}) {
	instance.Warnf(fmt.Sprintf("%s %s", header, template), args...)
}

func Errorf(header, template string, args ...interface{}) {
	instance.Errorf(fmt.Sprintf("%s %s", header, template), args...)
}

func Fatalf(header, template string, args ...interface{}) {
	instance.Fatalf(fmt.Sprintf("%s %s", header, template), args...)
}
