package logger

import (
	"log"
	"log/slog"
	"os"
)

type JsonLog struct {
	stdout *os.File
	file   *os.File
}

func NewJsonLog() Logger {
	file, err := os.Create("log.json")
	if err != nil {
		log.Println(err)
	}
	return &JsonLog{
		stdout: os.Stdout,
		file:   file,
	}
}

func (t *JsonLog) Write(p []byte) (n int, err error) {
	n, err = t.stdout.Write(p)
	if err != nil {
		return n, err
	}
	n, err = t.file.Write(p)
	return n, err
}

func (j *JsonLog) InfoLog(msg string, typeValue string, value any) {
	logger := slog.New(slog.NewJSONHandler(j, nil))
	logger.Info(msg, typeValue, value)
}

func (j *JsonLog) ErrorLog(msg string, err error) {
	logger := slog.New(slog.NewJSONHandler(j, nil))
	logger.Error(msg, "Error", err)
}

func (j *JsonLog) WarningLog(msg string) {
	logger := slog.New(slog.NewJSONHandler(j, nil))
	logger.Warn(msg)
}
