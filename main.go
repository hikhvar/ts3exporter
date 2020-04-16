package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/hikhvar/ts3exporter/pkg/collector"

	"github.com/hikhvar/ts3exporter/pkg/serverquery"
)

func main() {
	remote := flag.String("remote", "localhost:10011", "remote address of server query port")
	listen := flag.String("listen", ":9189", "listen address of the exporter")
	user := flag.String("user", "serveradmin", "the serverquery user of the ts3exporter")
	passwordFile := flag.String("passwordfile", "/etc/ts3exporter/password", "file containing the password. Must have 0600 permission. The file is not read if the environ")
	enableChannelMetrics := flag.Bool("enablechannelmetrics", false, "Enables the channel collector.")
	ignoreFloodLimits := flag.Bool("ignorefloodlimits", false, "Disable the server query flood limiter. Use this only if your exporter is whitelisted in the query_ip_whitelist.txt file.")

	flag.Parse()
	c, err := serverquery.NewClient(*remote, *user, mustReadPassword(*passwordFile), *ignoreFloodLimits)
	if err != nil {
		log.Fatalf("failed to init client %v\n", err)
	}
	sqInfo := serverquery.NewVirtualServer(c)
	sInfo := collector.NewServerInfo(sqInfo)
	mc := collector.NewMultiCollector(sInfo)

	if *enableChannelMetrics {
		cqInfo := serverquery.NewChannelView(c, sqInfo)
		cInfo := collector.NewChannel(cqInfo)
		mc.Add(cInfo)
	}

	prometheus.MustRegister(mc, collector.NewClient(c))
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listen, nil))
}

// mustReadPassword reads the password from the password file. The password file is only
// used if password is the empty string.
// If the file read fails, this function terminates the process.
func mustReadPassword(passwordFile string) string {
	fInfo, err := os.Stat(passwordFile)
	if err != nil {
		log.Fatalf("failed to get fileinfo of password file: %v\n", err)
	}
	if fInfo.Mode() != 0600 {
		log.Fatalf("password file permissions are to open. Have: %s, want: %o\n", fInfo.Mode().String(), 0600)
	}
	data, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		log.Fatalf("failed to read password file: %v\n", err)
	}

	// Trim possible line breaks that can be automatically added by e.g. vim
	return strings.Trim(string(data), "\r\n")
}
