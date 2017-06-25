package log

import (
	"github.com/ehlxr/go-utils/log"
	"github.com/sirupsen/logrus"
)

func init() {
	log.SetLogLevel(logrus.DebugLevel)
	// log.SetLogFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	// log.SetFn(false)
}

func TestLog() {

	log.Debug("debug text...")
	log.Info("info text...")
	log.Error("error text...")
	// log.Fatal("fatal text...")
	// log.Panic("panic text...")

	log.DebugWithFields(log.Fields{
		"id":   "test",
		"name": "jj",
	}, "debug with fields text...")

	log.InfoWithFields(log.Fields{
		"id":   "test",
		"name": "jj",
	}, "info with fields text...")

	log.ErrorWithFields(log.Fields{
		"id":   "test",
		"name": "jj",
	}, "error with fields text...")

	log.FatalWithFields(log.Fields{
		"id":   "test",
		"name": "jj",
	}, "fatal with fields text...")

	// log.Panic(log.Fields{
	// 	"id":   "test",
	// 	"name": "jj",
	// }, "fatal with fields text...")
}
