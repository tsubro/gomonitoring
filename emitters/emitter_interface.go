package emitters

import (
	"gomonitoring/context"
)

type Emitter interface {
	Emit(context *context.MonitoringContext)
}
