package emitters

type Level int

const (
	All Level = iota
	Debug
	Info
	Warn
	Error
	Off
	Trace
)

type LogLevel interface {
	GetLogLevel() Level
}

func (l Level) GetLogLevel() Level {
	return l
}
