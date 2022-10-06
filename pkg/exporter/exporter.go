package exporter

import (
	"net/http"

	"github.com/ShadowApex/meraki-exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run() {
	merakiCollector := collector.NewMerakiCollector()
	prometheus.MustRegister(merakiCollector)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9112", nil)
}
