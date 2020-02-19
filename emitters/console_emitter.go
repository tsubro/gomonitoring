package emitters

import (
	"gomonitoring/aggregators"
	"gomonitoring/context"
)

type ConsoleEmitterImpl struct {
	Level LogLevel
}

func (e ConsoleEmitterImpl) Emit(context *context.MonitoringContext, aggregator *aggregators.Aggregator) {
	
	logMetrics(context, e.Level)
	context.ClearLogs()
}
