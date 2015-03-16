package utils

import (
	"fmt"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/sirupsen/logrus"
)

var logr *logrus.Logger

type Level uint8

const (
	WarnLevel Level = iota
	InfoLevel
	DebugLevel
)

func ParseLevel(level string) (Level, error) {
	switch level {
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}
	var l Level
	return l, fmt.Errorf("not a valid logrus level: %q", level)
}

func getLogrusLevelDefault(def Level) (Level, error) {
	logLevelStr := configs.Get(configs.LOG_LEVEL)
	if logLevelStr == "" {
		return def, nil
	}

	logLevel, err := ParseLevel(logLevelStr)
	if err != nil {
		return 0, err
	}

	return logLevel, nil
}

func init() {
	if logr == nil {
		logr = logrus.New()
	}
}

func LogrusError(fname string, action string, err error) {
	logLevel, _ := getLogrusLevelDefault(DebugLevel)
	if logLevel != InfoLevel {
		logr.WithFields(logrus.Fields{
			"func_name": fname,
			"Action":    action,
		}).Errorf(err.Error())
	}
}

func LogrusInfo(fname, text string) {
	logLevel, _ := getLogrusLevelDefault(DebugLevel)
	if logLevel != WarnLevel {
		logr.WithFields(logrus.Fields{
			"func_name": fname,
		}).Info(text)
	}
}

func LogrusWarning(fname, action string, err error) {
	logLevel, _ := getLogrusLevelDefault(DebugLevel)
	if logLevel != InfoLevel {
		logr.WithFields(logrus.Fields{
			"func_name": fname,
			"Action":    action,
		}).Warn(err)
	}
}
