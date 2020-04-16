package collector

import "github.com/prometheus/client_golang/prometheus"

// ExporterMetrics tracks metrics internal to the exporter
type ExporterMetrics struct {
	refreshErrors *prometheus.CounterVec
}

func NewExporterMetrics() *ExporterMetrics {
	return &ExporterMetrics{
		refreshErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "exporter",
			Name:      "data_model_refresh_errors_total",
			Help:      "Errors encountered while updating the internal server model",
		}, []string{"collector"}),
	}
}

// RefreshError increases the refresh error counter of the given collector by one.
func (i *ExporterMetrics) RefreshError(collector string) {
	i.refreshErrors.WithLabelValues(collector).Inc()
}

func (i *ExporterMetrics) Describe(descs chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(i, descs)
}

func (i *ExporterMetrics) Collect(metrics chan<- prometheus.Metric) {
	i.refreshErrors.Collect(metrics)
}
