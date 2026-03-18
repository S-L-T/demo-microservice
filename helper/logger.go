package helper

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

type LogLevel uint8

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// Get the human-readable name of a log level
func (l LogLevel) String() string {
	n := []string{"panic", "fatal", "error", "warn", "info", "debug", "trace"}
	i := uint8(l)
	switch {
	case i <= uint8(TraceLevel):
		return n[i]
	case i > 0:
		return n[0]
	default:
		return strconv.Itoa(int(i))
	}
}

func InitializeLogger(minimumLevel LogLevel) error {
	logrus.SetReportCaller(true)
	l, err := logrus.ParseLevel(minimumLevel.String())
	if err != nil {
		return err
	}

	logrus.SetLevel(l)

	return nil
}

func Log(err error, level LogLevel) {
	switch level {
	case PanicLevel:
		logrus.Panic(err)
		break
	case FatalLevel:
		logrus.Fatal(err)
		break
	case ErrorLevel:
		logrus.Error(err)
		break
	case WarnLevel:
		logrus.Warn(err)
		break
	case InfoLevel:
		logrus.Info(err)
		break
	case DebugLevel:
		logrus.Debug(err)
		break
	case TraceLevel:
		logrus.Trace(err)
		break
	default:
		logrus.Warn("Incorrect logging level")
	}
}
