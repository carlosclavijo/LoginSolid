package logger

type Logger interface {
	InfoLog(msg string, typeValue string, value any)
	ErrorLog(msg string, err error)
	WarningLog(msg string)
}
