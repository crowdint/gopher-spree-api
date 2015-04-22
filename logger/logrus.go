package logger

import "github.com/Sirupsen/logrus"

type logrusLogger struct {
	*logrus.Logger
}

func NewLogrus() *logrusLogger {
	return &logrusLogger{logrus.New()}
}

func (l *logrusLogger) SetLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		l.Error(err)
	}
	l.Level = level
}
