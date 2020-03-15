package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/hikhvar/ts3exporter/pkg/collector"

	"github.com/hikhvar/ts3exporter/pkg/serverquery"
)

func main() {
	remote := flag.String("remote", "localhost:10011", "remote address of server query port")
	listen := flag.String("listen", ":9189", "listen address of the exporter")
	user := flag.String("user", "serveradmin", "the serverquery user of the ts3exporter")
	password := flag.String("password", "", "the serverquery password of the ts3exporter")
	flag.Parse()
	c, err := serverquery.NewClient(*remote, *user, *password)
	if err != nil {
		log.Fatalf("failed to init client %v\n", err)
	}
	sInfo := collector.NewServerInfo(c)
	mc := collector.NewMultiCollector(sInfo)

	prometheus.MustRegister(mc)
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listen, nil))
}
