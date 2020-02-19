package main

import (
	"gomonitoring/monitoring"
	"sync"

	"gomonitoring/emitters"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	log.Info("Test")

	e := emitters.FileEmitterImpl{Level : emitters.Info}
	m := monitoring.GetInstance()

	m.SetAggregatorFrequency(1)
	m.SetEmitterFrequency(1)
	m.SetEmitter(e)
	m.Start()

	m.PutMdc("JobId=5")

	m.Put("MyMain", "This is testing", "TAG1")
	m.Put("MyMain", "This is testing2", "TAG2")

	wg.Wait()

}
