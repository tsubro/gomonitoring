package monitoring

import (
	"gomonitoring/context"
	"gomonitoring/aggregators"
	"gomonitoring/emitters"
	"sync"
	"time"
)

type monitoring struct {
	aggregatorFrequency int //In minute
	emitterFrequency    int //In minute
	context             context.MonitoringContext
	aggregator          aggregators.Aggregator
	emitter             emitters.Emitter
}

var mux sync.Mutex
var obj *monitoring

func GetInstance() *monitoring {

	if obj == nil {
		mux.Lock()
		if obj == nil {
			obj = &monitoring{}
			obj.context = context.MonitoringContext{}
		}
		mux.Unlock()
	}
	return obj
}

func (m *monitoring) SetAggregatorFrequency(aggregatorFrequency int) {
	m.aggregatorFrequency = aggregatorFrequency
}

func (m *monitoring) SetEmitterFrequency(emitterFrequency int) {
	m.emitterFrequency = emitterFrequency
}

func (m *monitoring) GetContext() *context.MonitoringContext {
	return &m.context
}

func (m *monitoring) SetAggregator(aggregator aggregators.Aggregator) {
	m.aggregator = aggregator
}

func (m *monitoring) SetEmitter(emitter emitters.Emitter) {
	m.emitter = emitter
}

func (m *monitoring) Start() {
	if m.aggregator != nil {
		m.scheduleAggregator()
	}
	if m.emitter != nil {
		m.scheduleEmitter()
	}
}

func (m *monitoring) Stop() {
	m = nil
}

func (m *monitoring) scheduleAggregator() {
	ticker := time.NewTicker(time.Duration(m.aggregatorFrequency) * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				m.aggregator.Aggregate()
			}
		}
	}()
}

func (m *monitoring) scheduleEmitter() {
	ticker := time.NewTicker(time.Duration(m.emitterFrequency) * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				m.emitter.Emit(&m.context, &m.aggregator)
			}
		}
	}()
}

func (m *monitoring) Emit() {
	m.emitter.Emit(&m.context, &m.aggregator)
}

func (m *monitoring) Put(source string, logStr string, tag string) {
	m.context.PutLogs(source, logStr, tag)
}

func (m *monitoring) PutMdc(mdcStr string) {
	m.context.PutMdc(mdcStr)
}
