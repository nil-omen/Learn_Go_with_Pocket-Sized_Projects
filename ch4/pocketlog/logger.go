package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
	maxLength uint
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	l.logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	l.logf(LevelError, format, args...)
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
// There is no default maximum length - messages aren't trimmed.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
		maxLength: 0,
	}

	for _, configFunc := range opts {
		configFunc(lgr)
	}
	return lgr
}

// logf prints the message to the output. // Add decorations here, if any. #1
func (l *Logger) logf(lvl Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)

	if l.maxLength != 0 && uint(len([]rune(message))) > l.maxLength {
		message = string([]rune(message)[:l.maxLength]) + "[TRIMMED]"
	}
	_, _ = fmt.Fprintf(l.output, "%s %s\n", lvl, message)
}

// Logf formats and prints a message if the log level is high enough.
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}
	l.logf(lvl, format, args...)
}
