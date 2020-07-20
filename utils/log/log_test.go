package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSetting(t *testing.T) {
	SetFn(true)
	SetLogLevel(logrus.DebugLevel)
	SetLogFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	// SetLogFormatter(&prefixed.TextFormatter{
	// 	ForceFormatting:  true,
	// 	QuoteEmptyFields: true,
	// 	TimestampFormat:  "2006-01-02 15:04:05",
	// 	FullTimestamp:    true,
	// 	ForceColors:      true,
	// })
}

func TestInfo(t *testing.T) {
	Info("this is a Info log test.")

	InfoWithFields(Fields{
		"id":   10001,
		"name": "hello",
	}, "this is a InfoWithFields log test.")

	Infof("this is a %s log test.", "Infof")

	InfofWithFields(Fields{
		"id":   10001,
		"name": "hello",
	}, "this is a %s log test.", "InfofWithFields")
}

func TestPanic(t *testing.T) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recover:", x)
		}
	}()
	Panic("this is a Panic log test.")

	Panicf("this is a %s log test.", "Panicf")
}
