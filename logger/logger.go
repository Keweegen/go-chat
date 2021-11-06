package logger

import (
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

type Logger struct {
    native *zap.Logger
}

var logger *Logger

func GetLogger() *Logger {
    if logger != nil {
        return logger
    }

    infoLevel := zap.LevelEnablerFunc(func (lvl zapcore.Level) bool {
        return lvl == zapcore.InfoLevel || lvl == zapcore.WarnLevel
    })

    errLevel := zap.LevelEnablerFunc(func (lvl zapcore.Level) bool {
        return lvl == zapcore.ErrorLevel || lvl == zapcore.FatalLevel || lvl == zapcore.PanicLevel
    })

    core := zapcore.NewTee(
        zapcore.NewCore(
            zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
            zapcore.Lock(os.Stdout),
            infoLevel,
        ),
        zapcore.NewCore(
            zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
            zapcore.Lock(os.Stderr),
            errLevel,
        ),
    )

    logger = &Logger{}
    logger.native = zap.New(core)
    defer logger.native.Sync()

    return logger
}

func (l *Logger) With(args ...interface{}) {
    l.native.Sugar().With(args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
    l.native.Sugar().Infow(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
    l.native.Sugar().Errorw(msg, args...)
}
