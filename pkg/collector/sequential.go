package collector

import "github.com/prometheus/client_golang/prometheus"

// Ensure the collectors are run sequential. Since the different serverquery based collectors uses the same
// serverquery client, we need to run them sequential. Some serverquery command like `use` alternate the internal state.
type SequentialCollector []prometheus.Collector

func (s SequentialCollector) Describe(descs chan<- *prometheus.Desc) {
	for i := range s {
		s[i].Describe(descs)
	}
}

func (s SequentialCollector) Collect(metrics chan<- prometheus.Metric) {
	for i := range s {
		s[i].Collect(metrics)
	}
}
