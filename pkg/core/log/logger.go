package log

var DefaultLogger Logger

type Logger interface {
	GetOpt() Options

	Debug(msg string, keyvals ...interface{})
	Debugf(format string, val ...interface{})

	Info(msg string, keyvals ...interface{})
	Infof(format string, val ...interface{})

	Warn(msg string, keyvals ...interface{})
	Warnf(format string, val ...interface{})

	Error(msg string, keyvals ...interface{})
	Errorf(format string, val ...interface{})

	Fatalf(format string, val ...interface{})

	String() string
}

func New(l Logger) {
	DefaultLogger = l
}

func GetLogger() Logger {
	return DefaultLogger
}

func Info(msg string, keyvals ...interface{}) {
	DefaultLogger.Info(msg, keyvals...)
}

func Infof(msg string, val ...interface{}) {
	DefaultLogger.Infof(msg, val...)
}

func Debug(msg string, keyvals ...interface{}) {
	DefaultLogger.Debug(msg, keyvals...)
}

func Debugf(msg string, val ...interface{}) {
	DefaultLogger.Debugf(msg, val...)
}

func Warn(msg string, keyvals ...interface{}) {
	DefaultLogger.Warn(msg, keyvals...)
}

func Warnf(msg string, val ...interface{}) {
	DefaultLogger.Warnf(msg, val...)
}

func Error(msg string, keyvals ...interface{}) {
	DefaultLogger.Error(msg, keyvals...)
}

func Errorf(msg string, val ...interface{}) {
	DefaultLogger.Errorf(msg, val...)
}

func Fatal(msg string) {
	DefaultLogger.Fatalf(msg)
}

func Fatalf(msg string, val ...interface{}) {
	DefaultLogger.Fatalf(msg, val...)
}
