package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	verbose bool
	prefix  string
}

func New(verbose bool) *Logger {
	return &Logger{
		verbose: verbose,
		prefix:  "[nova-horizon]",
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.verbose {
		log.Printf("%s DEBUG: %s\n", l.prefix, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	log.Printf("%s INFO: %s\n", l.prefix, fmt.Sprintf(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s ERROR: %s\n", l.prefix, fmt.Sprintf(msg, args...))
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	fmt.Printf("%s WARN: %s\n", l.prefix, fmt.Sprintf(msg, args...))
}
