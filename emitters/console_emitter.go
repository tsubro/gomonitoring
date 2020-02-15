package emitters

import (
	"gomonitoring/context"
	"time"
	log "github.com/sirupsen/logrus"
)

type ConsoleEmitterImpl struct {
}

func (e ConsoleEmitterImpl) Emit(context *context.MonitoringContext) {

	log.Info(time.Now(), context.GetMdc(), context.GetLogs())
	context.ClearLogs()

}
