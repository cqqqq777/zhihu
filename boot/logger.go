package boot

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	g "zhihu/global"
)

func LoggerSetup() {
	encoder := GetEncoder()
	debugLv := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.DebugLevel
	})
	infoLv := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel
	})
	warnLv := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.WarnLevel
	})
	fatalLv := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.FatalLevel
	})
	var path = g.Config.Logger.SavePath
	core := [...]zapcore.Core{
		zapcore.NewCore(encoder, GetWriteSyncer(fmt.Sprintf("%v/all/all.log", path)), zapcore.DebugLevel),
		zapcore.NewCore(encoder, GetWriteSyncer(fmt.Sprintf("%v/debug/debug.log", path)), debugLv),
		zapcore.NewCore(encoder, GetWriteSyncer(fmt.Sprintf("%v/info/info.log", path)), infoLv),
		zapcore.NewCore(encoder, GetWriteSyncer(fmt.Sprintf("%v/warn/warn.log", path)), warnLv),
		zapcore.NewCore(encoder, GetWriteSyncer(fmt.Sprintf("%v/fatal/fatal.log", path)), fatalLv),
	}
	g.Logger = zap.New(zapcore.NewTee(core[:]...), zap.AddCaller())
	g.Logger.Info("init logger successfully")
}

func GetEncoder() zapcore.Encoder {
	var encoder zapcore.Encoder
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "caller",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeCaller:  zapcore.FullCallerEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		StacktraceKey: "stacktrace",
	}
	switch {
	case g.Config.Logger.LogType == "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case g.Config.Logger.LogType == "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	return encoder
}

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxAge:     g.Config.Logger.MaxAge,
		MaxBackups: g.Config.Logger.MaxGroups,
		MaxSize:    g.Config.Logger.MaxSize,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
