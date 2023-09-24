package logging

import (
	"log"
)

func init() {
	initLogging()
}

func initLogging() {
	logger := NewLogger()
	log.SetOutput(logger.Writer())
}
