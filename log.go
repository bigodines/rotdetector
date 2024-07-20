package rotdetector

import (
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

var (
	logger = log.New(os.Stdout, "rotdetector: ", log.Ldate|log.Ltime|log.Lshortfile)
	// logLevel is the current logging level.
	logLevel = INFO

	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

// SetLogLevel sets the logging level.
func SetLogLevel(level int) {
	logLevel = level
}

// Debug logs a message at the DEBUG level.
func Debug(v ...interface{}) {
	if logLevel <= DEBUG {
		logger.SetPrefix("DEBUG: ")
		logger.Printf(Red+" %v "+Reset, v...)
	}
}

// Info logs a message at the INFO level.
func Info(v ...interface{}) {
	if logLevel <= INFO {
		logger.SetPrefix("INFO: ")
		logger.Println(v...)
	}
}

// Warning logs a message at the WARNING level.
func Warning(v ...interface{}) {
	if logLevel <= WARNING {
		logger.SetPrefix("WARNING: ")
		logger.Println(v...)
	}
}

// Error logs a message at the ERROR level.
func Error(v ...interface{}) {
	if logLevel <= ERROR {
		logger.SetPrefix("ERROR: ")
		logger.Println(v...)
	}
}
