package log

type Event struct {
	Timestamp string      `json:"timestamp"`
	Level     Level       `json:"level"`
	Message   interface{} `json:"message"`
}

type Level string

const (
	TRACE Level = "trace"
	DEBUG Level = "debug"
	INFO  Level = "info"
	WARN  Level = "warn"
	ERROR Level = "error"
	FATAL Level = "fatal"
)
