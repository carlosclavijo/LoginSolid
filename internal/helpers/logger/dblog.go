package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/carlosclavijo/loginsolid/internal/repository"
)

type DbLog struct {
	R repository.Repository
}

func NewDbLog(r repository.Repository) Logger {
	return &DbLog{
		R: r,
	}
}

func (d *DbLog) InfoLog(msg string, typeValue string, value any) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(msg, typeValue, value)
	v := fmt.Sprintf("%+v", value)
	err := d.R.InsertLog("Info", msg, string(typeValue+": "+v))
	if err != nil {
		log.Println("Error", err)
	}
}

func (d *DbLog) ErrorLog(msg string, err error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Error(msg, "Error", err)
	e := d.R.InsertLog("Error", msg, "Error: "+err.Error())
	if e != nil {
		log.Println("Error", err)
	}
}

func (d *DbLog) WarningLog(msg string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Warn(msg)
	err := d.R.InsertLog("Warning", msg, "Entry denied")
	if err != nil {
		log.Println("Error", err)
	}
}
