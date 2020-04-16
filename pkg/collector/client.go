package collector

import (
	"github.com/hikhvar/ts3exporter/pkg/serverquery"
	"github.com/prometheus/client_golang/prometheus"
)

const clientSubsystem = "client"

// InstrumentedClient provides metrics from a serverquery client
type InstrumentedClient interface {
	Metrics() *serverquery.ClientMetrics
}

// Client is a collector providing metrics of the internal serverquery client
type Client struct {
	source InstrumentedClient

	Failed  *prometheus.Desc
	Success *prometheus.Desc
}

func NewClient(c InstrumentedClient) *Client {
	return &Client{
		source:  c,
		Failed:  prometheus.NewDesc(fqdn(clientSubsystem, "commands_failed_total"), "total failed server query command", nil, nil),
		Success: prometheus.NewDesc(fqdn(clientSubsystem, "commands_successful_total"), "total successful server query command", nil, nil),
	}
}

func (c *Client) Describe(desc chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, desc)
}

func (c *Client) Collect(met chan<- prometheus.Metric) {
	m := c.source.Metrics()
	met <- prometheus.MustNewConstMetric(c.Success, prometheus.CounterValue, float64(m.Success()))
	met <- prometheus.MustNewConstMetric(c.Failed, prometheus.CounterValue, float64(m.Failed()))
}
