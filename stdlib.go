package golog

import "log"

type goLogger struct {
	level Level
	log   *log.Logger
}

func NewGoLogger(level Level, log *log.Logger) Logger {
	return &goLogger{
		level: level,
		log:   log,
	}
}

func (l *goLogger) Print(level Level, v ...interface{}) {
	if l.level >= level {
		l.log.Print(v...)
	}
}

func (l *goLogger) Printf(level Level, format string, v ...interface{}) {
	if l.level >= level {
		l.log.Printf(format, v...)
	}
}
