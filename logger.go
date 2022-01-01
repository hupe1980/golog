package golog

type Logger interface {
	Print(level Level, v ...interface{})
	Println(level Level, v ...interface{})
	Printf(level Level, format string, v ...interface{})
}
