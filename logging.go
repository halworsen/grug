package grug

import (
	"fmt"
	"log"
)

const (
	logDebug = 1
	logInfo  = 2
	logError = 3
	logWarn  = 4
	logFatal = 5
)

var logLevelMap = map[int]string{
	1: "DEBUG",
	2: "INFO",
	3: "ERROR",
	4: "WARN",
	5: "FATAL",
}

// Log prints a log message tagged with "GRUG" and the log level
func (g *GrugSession) Log(level int, logMsg ...interface{}) {
	// Everything error and up is always logged
	// Everything is always logged before the config is loaded
	// DEBUG and INFO are logged with verbose logging enabled
	if level < logError && g.Config != nil && !g.Config.Verbose {
		return
	}
	logPrefix := fmt.Sprint("[GRUG/", logLevelMap[level], "]")

	var logFunc func(...interface{})
	switch level {
	case logDebug:
		logFunc = log.Println
	case logInfo:
		logFunc = log.Println
	case logError:
		logFunc = log.Println
	case logWarn:
		logFunc = log.Println
	case logFatal:
		logFunc = log.Fatalln
	}

	finalMsg := []interface{}{logPrefix}
	finalMsg = append(finalMsg, logMsg...)
	logFunc(finalMsg...)
}
