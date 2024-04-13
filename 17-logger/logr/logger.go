package logr

import (
	"os"

	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Log() *logrus.Logger {

	Logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
			ForceColors:     true,
		},
	}

	Logger.Out = ansicolor.NewAnsiColorWriter(os.Stdout)

	Logger.Info("Success message")
	Logger.Warning("Warning message")
	Logger.Error("Error message")

	return Logger
}
