package exporter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShadowApex/meraki-exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(port int) {
	listenPort := fmt.Sprintf(":%v", port)
	merakiCollector := collector.NewMerakiCollector()
	prometheus.MustRegister(merakiCollector)

	log.Println("Starting exporter on port", listenPort)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listenPort, nil)
}
