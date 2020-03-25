package collector

import (
	"github.com/hikhvar/ts3exporter/pkg/serverquery"
	"github.com/prometheus/client_golang/prometheus"
)

const channelSubsystem = "channel"

var channelLabels = []string{virtualServerLabel, channelLabel}

type Channel struct {
	channels ChannelInformer

	ClientsOnline *prometheus.Desc
	MaxClients    *prometheus.Desc
	Codec         *prometheus.Desc
	CodecQuality  *prometheus.Desc
	LatencyFactor *prometheus.Desc
	Unencrypted   *prometheus.Desc
	Permanent     *prometheus.Desc
	SemiPermanent *prometheus.Desc
	Default       *prometheus.Desc
	Password      *prometheus.Desc
}

// A ChannelInformer knows how to collect the data from all the channels on the monitoring target
type ChannelInformer interface {
	Refresh() error
	All() []serverquery.Channel
}

func NewChannel(ci ChannelInformer) *Channel {
	return &Channel{
		channels:      ci,
		ClientsOnline: prometheus.NewDesc(fqdn(channelSubsystem, "clients_online"), "number of clients currently online", channelLabels, nil),
		MaxClients:    prometheus.NewDesc(fqdn(channelSubsystem, "max_clients"), "maximal number of clients allowed in this channel", channelLabels, nil),
		Codec:         prometheus.NewDesc(fqdn(channelSubsystem, "codec"), "numeric number of configured codec for this channel", channelLabels, nil),
		CodecQuality:  prometheus.NewDesc(fqdn(channelSubsystem, "codec_quality"), "numeric number of codec quality level chosen for this channel", channelLabels, nil),
		LatencyFactor: prometheus.NewDesc(fqdn(channelSubsystem, "codec_latency_factor"), "numeric number of codec latency factor chosen for this channel", channelLabels, nil),
		Unencrypted:   prometheus.NewDesc(fqdn(channelSubsystem, "unencrypted"), "is the channel unencrypted", channelLabels, nil),
		Permanent:     prometheus.NewDesc(fqdn(channelSubsystem, "permanent"), "is the channel permanent", channelLabels, nil),
		SemiPermanent: prometheus.NewDesc(fqdn(channelSubsystem, "semi_permanent"), "is the channel semi permanent", channelLabels, nil),
		Default:       prometheus.NewDesc(fqdn(channelSubsystem, "default"), "is the channel the default channel", channelLabels, nil),
		Password:      prometheus.NewDesc(fqdn(channelSubsystem, "password"), "is the channel saved by an password", channelLabels, nil),
	}
}

func (c *Channel) Describe(desc chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, desc)
}

func (c *Channel) Collect(met chan<- prometheus.Metric) {
	for _, ch := range c.channels.All() {
		met <- prometheus.MustNewConstMetric(c.ClientsOnline, prometheus.GaugeValue, float64(ch.ClientsOnline), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.MaxClients, prometheus.GaugeValue, float64(ch.MaxClients), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.Codec, prometheus.GaugeValue, float64(ch.Codec), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.CodecQuality, prometheus.GaugeValue, float64(ch.CodecQuality), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.LatencyFactor, prometheus.GaugeValue, float64(ch.LatencyFactor), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.Unencrypted, prometheus.GaugeValue, float64(ch.Unencrypted), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.Permanent, prometheus.GaugeValue, float64(ch.Permanent), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.SemiPermanent, prometheus.GaugeValue, float64(ch.SemiPermanent), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.Default, prometheus.GaugeValue, float64(ch.Default), ch.HostingServer.Name, ch.Name)
		met <- prometheus.MustNewConstMetric(c.Password, prometheus.GaugeValue, float64(ch.Password), ch.HostingServer.Name, ch.Name)
	}
}

func (c *Channel) Refresh() error {
	return c.channels.Refresh()
}
