package logging

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new instance of DefaultLogger.
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}
}

func (l *Logger) Writer() io.Writer {
	return l.logger.Writer()
}

func (l *Logger) SetOutput(out io.Writer) {
	l.logger.SetOutput(out)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

func (l *Logger) Close() {}
