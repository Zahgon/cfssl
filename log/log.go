package log

const (
	LevelDebug = iota

	LevelInfo

	LevelWarning

	LevelError

	LevelCritical

	LevelFatal
)

var levelPrefix = [...]string{
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
	LevelWarning:  "WARNING",
	LevelError:    "ERROR",
	LevelCritical: "CRITICAL",
	LevelFatal:    "FATAL",
}

var Level = LevelInfo

type SyslogWriter interface {
	Debug(string)
	Info(string)
	Warning(string)
	Err(string)
	Crit(string)
	Emerg(string)
}

var syslogWriter SyslogWriter

func SetLogger(logger SyslogWriter) { _ = "STUB: not implemented"; return }

func print(l int, msg string) { _ = "STUB: not implemented"; return }

func outputf(l int, format string, v []interface{}) { _ = "STUB: not implemented"; return }

func output(l int, v []interface{}) { _ = "STUB: not implemented"; return }

func Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func Criticalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Critical(v ...interface{}) { _ = "STUB: not implemented"; return }

func Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func Warningf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warning(v ...interface{}) { _ = "STUB: not implemented"; return }

func Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debug(v ...interface{}) { _ = "STUB: not implemented"; return }
