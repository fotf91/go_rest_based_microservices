package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	// START configuration of zap using: config := zap.NewProductionConfig()
	config := zap.NewProductionConfig()                   // create configuration variable
	encoderConfig := zap.NewProductionEncoderConfig()     // create encoder
	encoderConfig.TimeKey = "timestamp"                   // define that the encoder will contain configuration of the "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // define the format of the timestamp
	config.EncoderConfig = encoderConfig                  // relate the configuration variable to the timestamp configuration
	log, err = config.Build(zap.AddCallerSkip(1))         // build using the above configuration
	/**
	CODE: zap.AddCallerSkip(1)
	EXPLANATION:
	contains zap.AddCallerSkip(1) which means that during the log that will be created
	the caller will not be the current file
	(the current file logger.go contains function Info that logs),
	but the file that called the current file
	so if main.go line 10 calls the logger.go function Info
	then in the log, the caller will be "main.go line 10"
	*/
	// END configuration of zap

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
