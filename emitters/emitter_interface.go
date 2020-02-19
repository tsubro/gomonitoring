package emitters

import (
	"fmt"
	"gomonitoring/aggregators"
	"gomonitoring/context"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = logrus.New()

type Emitter interface {
	Emit(context *context.MonitoringContext, aggregator *aggregators.Aggregator)
}

func getLogger() *logrus.Logger {
	log.SetLevel(logrus.DebugLevel)
	return log
}

func logMetrics(context *context.MonitoringContext, l LogLevel) {

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprint(time.Now()))
	builder.WriteString(fmt.Sprint(context.GetMdc()))
	builder.WriteString(fmt.Sprint(context.GetLogs()))
	s := builder.String()

	switch l {
	case All:
		log.Debug(s)
		log.Info(s)
		log.Warn(s)
		log.Error(s)
		log.Trace(s)
	case Debug:
		log.Debug(s)
	case Info:
		log.Info(s)
	case Warn:
		log.Warn(s)
	case Error:
		log.Error(s)
	case Trace:
		log.Trace(s)
	case Off:
	default:
		log.Info(s)
	}
}
