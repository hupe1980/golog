package golog

import (
	"errors"
	"strings"
)

type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}

// String returns the string representation of the logging level
func (p Level) String() string {
	return levelNames[p]
}

// LogLevel returns the log level from a string representation
func LogLevel(level string) (Level, error) {
	for i, name := range levelNames {
		if strings.EqualFold(name, level) {
			return Level(i), nil
		}
	}

	return ERROR, errors.New("invalid log level")
}
