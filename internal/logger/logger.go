package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitializeLogger() {
	isDebug := os.Getenv("DEBUG")
	logLevel := "info"

	if b, err := strconv.ParseBool(isDebug); err != nil {
		panic(err)
	} else if b {
		logLevel = "debug"
	}

	rawJSON := []byte(fmt.Sprintf(
		`{
		"level": "%s",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoding": "json",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`, strings.ToLower(logLevel),
	))

	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	var err error

	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	var zapLog *zap.Logger
	zapLog, err = cfg.Build()
	zap.ReplaceGlobals(zapLog)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = zapLog.Sync()
	}()
	zapLog.Info(
		fmt.Sprintf(
			"logger construction succeeded, loglevel: %s", logLevel,
		),
	)
}
