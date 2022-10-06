package collector

import "github.com/prometheus/client_golang/prometheus"

type Settable interface {
	Set(ch chan<- prometheus.Metric, value float64, labels ...string)
}

type Metric struct {
	desc *prometheus.Desc
}

func NewMetric(desc *prometheus.Desc) *Metric {
	return &Metric{desc: desc}
}

type Counter struct {
	Metric
}

func (c *Counter) Set(ch chan<- prometheus.Metric, value float64, labels ...string) {
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, value, labels...)
}

func NewCounter(desc *prometheus.Desc) *Counter {
	return &Counter{Metric: Metric{desc: desc}}
}

type Gauge struct {
	Metric
}

func (g *Gauge) Set(ch chan<- prometheus.Metric, value float64, labels ...string) {
	ch <- prometheus.MustNewConstMetric(g.desc, prometheus.GaugeValue, value, labels...)
}

func NewGauge(desc *prometheus.Desc) *Gauge {
	return &Gauge{Metric: Metric{desc: desc}}
}
