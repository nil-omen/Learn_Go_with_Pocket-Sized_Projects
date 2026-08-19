package pocketlog

import "io"

// Option defines a functional option to our logger.
type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}

// // WithMaxLength sets the maximum length, in characters, of a message.
// Use 0 for no maximum length.
func WithMaxLength(maxLength uint) Option {
	return func(lgr *Logger) {
		lgr.maxLength = maxLength
	}
}
