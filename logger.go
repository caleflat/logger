package logger

import (
	"fmt"
	"io"
	"os"
)

type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	}

	panic("unreachable")
}

// Logger is a simple logger with level, color and timestamp.
// It's thread-safe.
type Logger struct {
	File            io.Writer
	TimestampFormat string
	Level           Level
	UseColor        bool
	UseTimestamp    bool
}

func New() *Logger {
	return &Logger{
		Level:           LevelTrace,
		UseColor:        true,
		File:            os.Stdout,
		UseTimestamp:    true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

func (l *Logger) Trace(format string, v ...interface{}) {
	if l.Level <= LevelTrace {
		l.print(LevelTrace, format, v...)
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.Level <= LevelDebug {
		l.print(LevelDebug, format, v...)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.Level <= LevelInfo {
		l.print(LevelInfo, format, v...)
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.Level <= LevelWarn {
		l.print(LevelWarn, format, v...)
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.Level <= LevelError {
		l.print(LevelError, format, v...)
	}
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	if l.Level <= LevelFatal {
		l.print(LevelFatal, format, v...)
	}
}

func (l *Logger) print(level Level, format string, v ...interface{}) {
	originalFormat := format
	if l.UseTimestamp {
		format = fmt.Sprintf("%s [%s] ", l.TimestampFormat, level.String())
	}
	if l.UseColor {
		switch level {
		case LevelTrace:
			format = "\x1b[34m" + format + "\x1b[0m"
		case LevelDebug:
			format = "\x1b[93m" + format + "\x1b[0m"
		case LevelInfo:
			format = "\x1b[92m" + format + "\x1b[0m"
		case LevelWarn:
			format = "\x1b[33m" + format + "\x1b[0m"
		case LevelError:
			format = "\x1b[31m" + format + "\x1b[0m"
		case LevelFatal:
			format = "\x1b[35m" + format + "\x1b[0m"
		}
	}
	format += originalFormat + "\n"
	_, _ = fmt.Fprintf(l.File, format, v...)
}
