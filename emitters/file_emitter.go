package emitters

import (
	"gomonitoring/aggregators"
	"gomonitoring/context"
	"os"
	"strings"
)

type FileEmitterImpl struct {
	LogFileName string
	Level       LogLevel
}

func (e FileEmitterImpl) Emit(context *context.MonitoringContext, aggregator *aggregators.Aggregator) {

	log := getLogger()
	
	logFileName := e.LogFileName
	if logFileName == "" {
		logFileName = "temp.log"
	} else if !strings.HasSuffix(logFileName, ".log") && !strings.HasSuffix(logFileName, ".txt") {
		logFileName = e.LogFileName + ".log"
	}

	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Error("Failed to Open File to Log")
	}
	defer f.Close()

	log.SetOutput(f)
	logMetrics(context, e.Level)
	context.ClearLogs()
}
