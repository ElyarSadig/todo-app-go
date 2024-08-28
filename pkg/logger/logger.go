package logger

type Logger interface {
	Info(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Debug(v ...any)
	Fatal(v ...any)
}
