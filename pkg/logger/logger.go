package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLoggerFile(path string) *os.File {
	logsFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		logrus.Error(err)
	}

	logsOut := io.MultiWriter(os.Stdout, logsFile)

	logrus.SetOutput(logsOut)
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})

	return logsFile
}
