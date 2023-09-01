package logger

import (
	"dp_test/shared"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const dateTimeFormat = "02/01/2006:15:04:05"

var defaultLogger *zap.SugaredLogger
var logFilesObjRef *os.File

func init() {
	defaultLogger = getLogger(zap.InfoLevel, "log.txt", true)
}

func WithWrappedModel(sm shared.Wrapped) *zap.SugaredLogger {
	return defaultLogger.With("Model", sm)
}

// Gets Zap logger with specified instance
//
// returns a SugaredLogger instance
func getLogger(level zapcore.Level, filename string, logSTDIO bool) *zap.SugaredLogger {
	var cores []zapcore.Core

	// Configure logger options
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(dateTimeFormat))
	}

	logFile, errLogFile := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if errLogFile != nil {
		defer logFile.Close()
		panic("Failed to open log file " + filename)
	}

	var err error
	fileCore := zapcore.NewCore(zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(logFile), level)

	cores = append(cores, fileCore)

	if logSTDIO {
		// Add stdio output core
		stdoutCore := zapcore.NewCore(zapcore.NewConsoleEncoder(config),
			zapcore.Lock(os.Stdout), level)
		cores = append(cores, stdoutCore)
	}

	core := zapcore.NewTee(cores...)

	_logger := zap.New(core).Sugar()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	logFilesObjRef = logFile
	return _logger
}

func CloseLogFile() {
	if logFilesObjRef != nil {
		logFilesObjRef.Close()
	}
}
