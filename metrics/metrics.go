package metrics

import (
	"sync"
	"time"

	mt "github.com/rcrowley/go-metrics"
)

type MeticsTimers struct {
	timers map[string]mt.Timer
	mu     sync.RWMutex
}

func (m *MeticsTimers) Measure(elapsed time.Duration, key string) {
	m.mu.RLock()
	t, ok := m.timers[key]
	m.mu.RUnlock()

	if !ok {
		m.mu.Lock()
		t, ok = m.timers[key]
		if !ok {
			t = mt.NewTimer()
			m.timers[key] = t
		}
		m.mu.Unlock()
	}

	t.Update(elapsed)
}

func (m *MeticsTimers) Get() map[string]map[string]float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make(map[string]map[string]float64)
	for query, timer := range m.timers {
		result[query] = map[string]float64{
			"count": float64(timer.Count()),
			"min":   float64(timer.Min()) / float64(time.Millisecond),
			"max":   float64(timer.Max()) / float64(time.Millisecond),
			"avg":   timer.Mean() / float64(time.Millisecond),
			"rate":  timer.Rate1(),
			"p50":   timer.Percentile(0.5) / float64(time.Millisecond),
			"p75":   timer.Percentile(0.75) / float64(time.Millisecond),
			"p95":   timer.Percentile(0.95) / float64(time.Millisecond),
			"p99":   timer.Percentile(0.99) / float64(time.Millisecond),
		}
	}
	return result
}

type MetricsDB struct {
	connections mt.Histogram
	queries     mt.Meter
	executes    mt.Meter
	rows        mt.Meter
	affects     mt.Meter
	timers      *MeticsTimers
}

func NewMetricsDB() *MetricsDB {
	return &MetricsDB{
		connections: mt.NewHistogram(mt.NewExpDecaySample(1028, 0.015)),
		queries:     mt.NewMeter(),
		executes:    mt.NewMeter(),
		rows:        mt.NewMeter(),
		affects:     mt.NewMeter(),
		timers:      &MeticsTimers{timers: make(map[string]mt.Timer)},
	}
}

func (m *MetricsDB) Timers() *MeticsTimers {
	return m.timers
}

func (m *MetricsDB) MarkQueries(v int) {
	if v != 0 {
		m.queries.Mark(int64(v))
	}
}

func (m *MetricsDB) MarkExecutes(v int) {
	if v != 0 {
		m.executes.Mark(int64(v))
	}
}

func (m *MetricsDB) MarkRows(v int) {
	if v != 0 {
		m.rows.Mark(int64(v))
	}
}

func (m *MetricsDB) MarkAffects(v int) {
	if v != 0 {
		m.affects.Mark(int64(v))
	}
}

func (m *MetricsDB) MarkConnections(v int) {
	m.connections.Update(int64(v))
}

func (m *MetricsDB) Measure(start time.Time, query string) {
	m.timers.Measure(time.Now().Sub(start), query)
}

func (m *MetricsDB) Get() map[string]float64 {
	return map[string]float64{
		"connections_min": float64(m.connections.Min()),
		"connections_max": float64(m.connections.Max()),
		"connections_avg": m.connections.Mean(),
		"queries_count":   float64(m.queries.Count()),
		"queries_rate":    m.queries.Rate1(),
		"executes_count":  float64(m.executes.Count()),
		"executes_rate":   m.executes.Rate1(),
		"rows_count":      float64(m.rows.Count()),
		"rows_rate":       m.rows.Rate1(),
		"affects_count":   float64(m.affects.Count()),
		"affects_rate":    m.affects.Rate1(),
	}
}

type MetricsHttp struct {
	clients        mt.Histogram
	clientsCounter mt.Counter
	requests       mt.Timer
	status2xx      mt.Meter
	status3xx      mt.Meter
	status4xx      mt.Meter
	status5xx      mt.Meter
	statusErr      mt.Meter
	timers         *MeticsTimers
}

func NewMetricsHttp() *MetricsHttp {
	return &MetricsHttp{
		clients:        mt.NewHistogram(mt.NewExpDecaySample(1028, 0.015)),
		clientsCounter: mt.NewCounter(),
		requests:       mt.NewTimer(),
		status2xx:      mt.NewMeter(),
		status3xx:      mt.NewMeter(),
		status4xx:      mt.NewMeter(),
		status5xx:      mt.NewMeter(),
		statusErr:      mt.NewMeter(),
		timers:         &MeticsTimers{timers: make(map[string]mt.Timer)},
	}
}

func (m *MetricsHttp) Timers() *MeticsTimers {
	return m.timers
}

func (m *MetricsHttp) IncClient() {
	m.clientsCounter.Inc(1)
}

func (m *MetricsHttp) DecClient() {
	m.clientsCounter.Dec(1)
}

func (m *MetricsHttp) Mark2xx() {
	m.status2xx.Mark(1)
}

func (m *MetricsHttp) Mark3xx() {
	m.status3xx.Mark(1)
}

func (m *MetricsHttp) Mark4xx() {
	m.status4xx.Mark(1)
}

func (m *MetricsHttp) Mark5xx() {
	m.status5xx.Mark(1)
}

func (m *MetricsHttp) MarkErr() {
	m.statusErr.Mark(1)
}

func (m *MetricsHttp) Measure(elapsed time.Duration, uri string) {
	m.requests.Update(elapsed)
	m.timers.Measure(elapsed, uri)
}

func (m *MetricsHttp) Get() map[string]float64 {
	m.clients.Update(m.clientsCounter.Count())
	return map[string]float64{
		"clients_min":    float64(m.clients.Min()),
		"clients_max":    float64(m.clients.Max()),
		"clients_avg":    m.clients.Mean(),
		"requests_count": float64(m.requests.Count()),
		"request_rate":   m.requests.Rate1(),
		"times_min":      float64(m.requests.Min()) / float64(time.Millisecond),
		"times_max":      float64(m.requests.Max()) / float64(time.Millisecond),
		"times_avg":      m.requests.Mean() / float64(time.Millisecond),
		"times_50":       m.requests.Percentile(0.5) / float64(time.Millisecond),
		"times_75":       m.requests.Percentile(0.75) / float64(time.Millisecond),
		"times_95":       m.requests.Percentile(0.95) / float64(time.Millisecond),
		"times_99":       m.requests.Percentile(0.99) / float64(time.Millisecond),
		"2xx_count":      float64(m.status2xx.Count()),
		"2xx_rate":       m.status2xx.Rate1(),
		"3xx_count":      float64(m.status3xx.Count()),
		"3xx_rate":       m.status3xx.Rate1(),
		"4xx_count":      float64(m.status4xx.Count()),
		"4xx_rate":       m.status4xx.Rate1(),
		"5xx_count":      float64(m.status5xx.Count()),
		"5xx_rate":       m.status5xx.Rate1(),
		"err_count":      float64(m.statusErr.Count()),
		"err_rate":       m.statusErr.Rate1(),
	}
}
