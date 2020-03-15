package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type TS3Collector interface {
	prometheus.Collector
	Refresh() error
}

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
