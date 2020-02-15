package context

import( 
	"time"
	"sync"
)

type MonitoringContext struct {
	sync.RWMutex
	mdc        []string
	logContext []logContext
}

func (m MonitoringContext) GetMdc() []string {
	return m.mdc
}

func (m *MonitoringContext) PutMdc(l string) {
	m.Lock()
	defer m.Unlock()
	m.mdc = append(m.mdc, l)
}

func (m *MonitoringContext) ClearMdc() {
	m.Lock()
	defer m.Unlock()
	m.mdc = nil
}

func (m MonitoringContext) GetLogs() []logContext {
	return m.logContext
}

func (m *MonitoringContext) PutLogs(source string, logStr string, tag string) {
	logContextObj := logContext{}
	logContextObj.logTime = time.Now()
	logContextObj.Source = source
	logContextObj.logStr = logStr
	logContextObj.tag = tag

	m.Lock()
	defer m.Unlock()
	m.logContext = append(m.logContext, logContextObj)
}

func (m *MonitoringContext) ClearLogs() {
	m.Lock()
	defer m.Unlock()
	m.logContext = nil
}

type logContext struct {
	logTime time.Time
	Source  string
	logStr  string
	tag     string
}
