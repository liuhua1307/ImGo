package log

var log Logger

func RegisterLog(logger Logger) {
	log = logger
}

func Log() Logger {
	if log == nil {
		panic("implement not found for interface Logger, please register")
	}
	return log
}

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
}

type Field struct {
	Key   string
	Value interface{}
}
