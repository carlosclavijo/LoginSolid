package logger

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type TcpLog struct {
	LogValues []byte `json:"log_value"`
}

func (t *TcpLog) InfoLog(msg string, typeValue string, value any) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(msg, typeValue, value)
	js := fmt.Sprintf("%+v", value)
	t.LogValues = []byte(js)
	resp := logRequest(t.LogValues)
	defer resp.Body.Close()
}

func (t *TcpLog) ErrorLog(msg string, err error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Error(msg, "Error", err)
	js := fmt.Sprintf("%+v", msg)
	t.LogValues = []byte(js)
	resp := logRequest(t.LogValues)
	defer resp.Body.Close()
}

func (t *TcpLog) WarningLog(msg string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Warn(msg)
	js := fmt.Sprintf("%+v", msg)
	t.LogValues = []byte(js)
	resp := logRequest(t.LogValues)
	defer resp.Body.Close()
}

func logRequest(js []byte) *http.Response {
	req, err := http.NewRequest("POST", "https://www.sd-bo.com/log.php?log_value=hello", bytes.NewBuffer(js))
	if err != nil {
		log.Println(err)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp
}
