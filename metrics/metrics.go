package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MemCfgStored = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "scfg",
		Name: "mem_cfg_stored_total",
		Help: "The number of stored config sets in memory cache",
	})

	ResponseTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "scfg",
		Name: "response_time_milliseconds",
		Help: "The response time of endpoint",
	}, []string{"handler"})
)

func init() {
	prometheus.DefaultRegisterer.MustRegister(MemCfgStored, ResponseTime)
}