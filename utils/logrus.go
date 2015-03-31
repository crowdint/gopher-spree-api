package utils

import (
	"runtime"

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

func getLogrusLevelDefault(def Level) Level {
	levels := map[string]Level{
		"warn":  WarnLevel,
		"info":  InfoLevel,
		"debug": DebugLevel,
	}

	logLevel, isLevel := levels[configs.Get(configs.LOG_LEVEL)]

	if !isLevel {
		return def
	}
	return logLevel
}

func init() {
	if logr == nil {
		logr = logrus.New()
	}
}

func LogrusError(fname string, err error) {
	logLevel := getLogrusLevelDefault(DebugLevel)
	if logLevel != InfoLevel {
		logr.WithFields(logrus.Fields{
			"File name": fname,
		}).Errorf(err.Error())
	}
}

func LogrusInfo(fname, text string) {
	logLevel := getLogrusLevelDefault(DebugLevel)
	if logLevel != WarnLevel {
		logr.WithFields(logrus.Fields{
			"File name": fname,
		}).Info(text)
	}
}

func LogrusWarning(fname string, err error) {
	logLevel := getLogrusLevelDefault(DebugLevel)
	if logLevel != InfoLevel {
		logr.WithFields(logrus.Fields{
			"File name": fname,
		}).Warn(err)
	}
}

func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
