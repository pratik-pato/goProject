package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota + 1
)

//std = New(os.Stderr, "", LstdFlags)
var (
	l        *log.Logger // PRIVATE LOGGER INSTANCE
	logLevel LogLevel    // PRIVATE LOGGING LEVEL
)

func callDebug() {
	Debug("Debug message for %s", "pratik")
}

type LoggerConfig struct {
	logFile  string
	logLevel int
}

func llog(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	//std.Output(2, fmt.Sprintf(format, v...)
	llog(format, v)
	//TODO handle log level
}

func SetLogLevel(level LogLevel) {
	//TODO
}

func Init(config *LoggerConfig) bool {
	//TODO file hndlin
	fileHandle := os.Stderr
	l = log.New(fileHandle, "", log.Ldate|log.Ltime|log.Lshortfile)
	l.Println("Logger initialized successfully.")
	return true
}
