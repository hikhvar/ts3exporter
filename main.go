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
	password := flag.String("password", "", "the serverquery password of the ts3exporter")
	passwordFile := flag.String("passwordfile", "/etc/ts3exporter/password", "file containing the password. Only read if -password not set. Must have 0600 permission.")

	flag.Parse()
	c, err := serverquery.NewClient(*remote, *user, mustReadPassword(*password, *passwordFile))
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

// mustReadPassword reads the password either from the commandline flag or the password file. The password file is only
// used if password is the empty string.
// If the file read fails, this function terminates the process.
func mustReadPassword(password, passwordFile string) string {
	if password != "" {
		return password
	}
	fInfo, err := os.Stat(passwordFile)
	if err != nil {
		log.Fatalf("failed to get fileinfo of password file: %v\n", err)
	}
	if fInfo.Mode() != 0600 {
		log.Fatalf("password file permissions are to open. Have: %o, want: %o\n", fInfo.Mode(), 0600)
	}
	data, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		log.Fatalf("failed to read password file: %v\n", err)
	}

	// Trim possible line breaks that can be automatically added by e.g. vim
	return strings.Trim(string(data), "\r\n")
}
