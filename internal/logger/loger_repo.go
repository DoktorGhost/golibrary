package logger

// Logger - интерфейс для логирования
type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}
