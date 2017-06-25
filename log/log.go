package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

type (
	Fields logrus.Fields
)

var (
	mylog *logrus.Logger
	fn    bool
)

func init() {
	mylog = logrus.New()
	mylog.Formatter = &logrus.JSONFormatter{}
	mylog.Out = os.Stdout
	mylog.Level = logrus.DebugLevel

	mylog.Formatter = &prefixed.TextFormatter{
		ForceFormatting:  true,
		QuoteEmptyFields: true,
		TimestampFormat:  "2006-01-02 15:04:05",
		FullTimestamp:    true,
		ForceColors:      true,
	}

	fn = true
}

func SetFn(val bool) {
	fn = val
}

func SetLogLevel(level logrus.Level) {
	mylog.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	mylog.Formatter = formatter
}

func NewLog(prefix string) *logrus.Entry {
	return mylog.WithField("prefix", prefix)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func Debug(args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Debug(args...)
	} else {
		mylog.Debug(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Debugf(format, args...)
	} else {
		mylog.Debugf(format, args...)
	}
}

func DebugWithFields(fields Fields, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Debug(args...)
}

func DebugfWithFields(fields Fields, format string, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Debugf(format, args...)
}

func Info(args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Info(args...)
	} else {
		mylog.Info(args...)
	}
}

func Infof(format string, args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Infof(format, args...)
	} else {
		mylog.Infof(format, args...)
	}
}

func InfoWithFields(fields Fields, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Info(args...)

}

func InfofWithFields(fields Fields, format string, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Infof(format, args...)
}

func Error(args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Error(args...)
	} else {
		mylog.Error(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Errorf(format, args...)
	} else {
		mylog.Errorf(format, args...)
	}
}

func ErrorWithFields(fields Fields, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Error(args...)

}

func ErrorfWithFields(fields Fields, format string, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Fatal(args...)
	} else {
		mylog.Fatal(args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Fatalf(format, args...)
	} else {
		mylog.Fatalf(format, args...)
	}
}

func FatalWithFields(fields Fields, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Fatal(args...)
}

func FatalfWithFields(fields Fields, format string, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Panic(args...)
	} else {
		mylog.Panic(args...)
	}
}

func Panicf(format string, args ...interface{}) {
	if fn {
		mylog.WithField("prefix", fileInfo(2)).Panicf(format, args...)
	} else {
		mylog.Panicf(format, args...)
	}
}

func PanicWithFields(fields Fields, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Panic(args...)
}

func PanicfWithFields(fields Fields, format string, args ...interface{}) {
	entry := mylog.WithFields(logrus.Fields(fields))

	if fn {
		entry.Data["prefix"] = fileInfo(2)
	}
	entry.Panicf(format, args...)
}
