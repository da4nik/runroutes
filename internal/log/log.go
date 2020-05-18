package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Log struct {
	log *logrus.Logger
}

type Logger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

func NewLogger(filename string) *Log {
	var out io.Writer
	out = os.Stdout

	logFile, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0755)
	if err == nil {
		out = io.MultiWriter(os.Stdout, logFile)
	}

	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	log.SetOutput(out)

	return &Log{
		log: log,
	}
}

func (l *Log) Infof(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	l.log.Debugf(format, v...)
}

func (l *Log) Errorf(format string, v ...interface{}) {
	l.log.Errorf(format, v...)
}

func (l *Log) WithModule(module string) Logger {
	newLog := l.log.WithField("module", module)

	return newLog
}
