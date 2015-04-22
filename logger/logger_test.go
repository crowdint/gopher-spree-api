package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestDebug(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("debug")

	Debug("foo")

	if strings.Contains("foo", out.String()) {
		t.Error("a debug message was expected")
	}
}

func TestInfo(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("info")

	Info("foo")

	if strings.Contains("foo", out.String()) {
		t.Error("an info message was expected")
	}
}

func TestWarn(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("warn")

	Warn("foo")

	if strings.Contains("foo", out.String()) {
		t.Error("a warn message was expected")
	}
}

func TestError(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("error")

	Error("foo")

	if strings.Contains("foo", out.String()) {
		t.Error("an error message was expected")
	}
}

func TestDebugf(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("debug")

	Debugf("foo %s", "bar")

	if strings.Contains("foo bar", out.String()) {
		t.Error("a debug message was expected")
	}
}

func TestInfof(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("info")

	Infof("foo %s", "bar")

	if strings.Contains("foo bar", out.String()) {
		t.Error("an info message was expected")
	}
}

func TestWarnf(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("warn")

	Warnf("foo %s", "bar")

	if strings.Contains("foo bar", out.String()) {
		t.Error("a warn message was expected")
	}
}

func TestErrorf(t *testing.T) {
	out := bytes.NewBufferString("")

	logger := NewLogrus()
	logger.Out = out

	SetLogger(logger)
	SetLevel("error")

	Errorf("foo %s", "bar")

	if strings.Contains("foo bar", out.String()) {
		t.Error("an error message was expected")
	}
}
