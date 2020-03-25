package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type TS3Collector interface {
	prometheus.Collector
	Refresh() error
}

// A MultiCollector implements the prometheus.Collector interface. The MultiCollector triggers a Refresh on all managed
// TS3Collectors before calling their Collect function. Thus every TS3Collector should provide the current values upon
// scrape.
type MultiCollector struct {
	collectors    []TS3Collector
	refreshErrors prometheus.Counter
}

func NewMultiCollector(collectors ...TS3Collector) *MultiCollector {
	return &MultiCollector{
		collectors: collectors,
		refreshErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "exporter",
			Name:      "data_model_refresh_errors_total",
			Help:      "Errors encountered while updating the internal server model",
		}),
	}
}

// Add adds a new TS3Collector to the collectors managed by this MultiCollector
func (m *MultiCollector) Add(c TS3Collector) {
	m.collectors = append(m.collectors, c)
}

func (m *MultiCollector) Describe(c chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(m, c)
}

func (m *MultiCollector) Collect(c chan<- prometheus.Metric) {
	for _, col := range m.collectors {
		err := col.Refresh()
		if err != nil {
			m.refreshErrors.Inc()
			log.Printf("failed to refresh collector: %v", err)
		}
	}
	m.refreshErrors.Collect(c)
	for _, col := range m.collectors {
		col.Collect(c)
	}
}
