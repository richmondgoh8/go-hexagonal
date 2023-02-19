package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var zapSugarLog *zap.SugaredLogger

func Init() {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	var globalLoggingLevel zapcore.Level

	switch os.Getenv("LOGGING_LEVEL") {
	case "DEBUG":
		globalLoggingLevel = zapcore.DebugLevel
	case "INFO":
		globalLoggingLevel = zapcore.InfoLevel
	case "WARN":
		globalLoggingLevel = zapcore.WarnLevel
	case "ERROR":
		globalLoggingLevel = zapcore.ErrorLevel
	default:
		globalLoggingLevel = zapcore.DebugLevel
		log.Printf("Warning! missing or invalid LOGGING_LEVEL ENV Variable, Setting Default Logging Level to %s\n", globalLoggingLevel)
	}

	config.Level = zap.NewAtomicLevelAt(globalLoggingLevel)

	value, exists := os.LookupEnv("SERVICE_NAME")
	if exists {
		config.InitialFields = map[string]interface{}{
			"engine": value,
		}
	} else {
		config.InitialFields = map[string]interface{}{
			"engine": "not_specified",
		}
	}

	// Caller Skip allows the proper caller instead of showing pkg caller
	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	// clears buffers if any
	defer zapLogger.Sync()
	zapSugarLog = zapLogger.Sugar()
}

func Debug(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Debugw(msg, keyValueField...)
}

func Info(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Infow(msg, keyValueField...)
}

func Warn(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Warnw(msg, keyValueField...)
}

func Error(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Errorf(msg, keyValueField...)
}

func getTrackingDetails(ctx context.Context, keyValueField []interface{}) []interface{} {
	if val, exist := ctx.Value("trackingID").(string); exist {
		keyValueField = append(keyValueField, "tracking_id", val)
	}

	return keyValueField
}
