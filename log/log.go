package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
	"time"
)

var sugaredLogger *zap.SugaredLogger
var atomicLevel zap.AtomicLevel
var ioWrite *lumberjack.Logger

func init() {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	// define default level as debug level
	atomicLevel = zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	//core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, atomicLevel)
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), getLogWriter(), atomicLevel)
	logger := zap.New(core)
	logger = logger.WithOptions(zap.AddCallerSkip(1))
	logger = logger.WithOptions(zap.AddCaller())
	sugaredLogger = logger.Sugar()
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05.000"))
}

func Logger() *zap.Logger {
	return sugaredLogger.Desugar()
}

func SetLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

func Fatal(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	sugaredLogger.Panicf(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}

func Rotate() {
	fmt.Println("logger.Rotate....")
	ioWrite.Rotate()
}

func getLogWriter() zapcore.WriteSyncer {
	fileName := fmt.Sprintf("%d-%d-%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	//fileName := fmt.Sprintf("%s.log", pCfg.Name)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path.Join("./", fileName),
		MaxSize:    1024, //MB
		MaxBackups: 0,    //文件数量 -0 无限制
		MaxAge:     22,   //保留时间 -0 永久保存
		Compress:   true, //压缩
	}
	ioWrite = lumberJackLogger
	return zapcore.AddSync(ioWrite)
}
