package main

import (
	"flag"
	"log"
	"net/http"

	"mirakurun-exporter/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen address", ":9100", "The Address to listen on for HTTP Requests.")
)

func main() {
	flag.Parse()

	s := collector.NewStatusCollector()
	v := collector.NewVersionCollector()

	reg := prometheus.NewRegistry()
	reg.MustRegister(s, v)

	http.HandleFunc("/", indexPage)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	log.Println("Listening on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Mirakurun Exporter</title>
		</head>
		<body>
			<h1>Mirakurn Exporter</h1>
			<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>
	`))
}
