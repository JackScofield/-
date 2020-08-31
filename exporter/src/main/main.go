package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("listen-address",":8080","The address to listen on for HTTP request" )
var (
	opsQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "apu_studio",
		Subsystem: "blob_storage",
		Name: "ops_queued",
		Help: "number of blog storage operation writing to be processed.",
	})
	taskerCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "apu_studio",
		Subsystem: "worker_pool",
		Name: "completed_tasks_total",
		Help: "Total number of tasks completed.",
	})
	temps = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "pond_temperature_celsius",
		Help: "The temperature of the frog pond",
		Objectives: map[float64]float64{0.5:0.05,0.9:0.01,0.99:0.001},
	})

	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "rpc_durations_histogram_seconds",
		Help: "RPC latency distributions.",
		Buckets: []float64{1,2,5,10,20,60},
	})
)


func init() {
	prometheus.MustRegister(opsQueued)
	prometheus.MustRegister(taskerCounter)
	prometheus.MustRegister(temps)
	prometheus.MustRegister(rpcDurationsHistogram)
}

func main() {
	flag.Parse()
	go func() {
		for {
			opsQueued.Add(4)
			taskerCounter.Inc()
			temps.Observe(4)
			
			time.Sleep(time.Second * 1)
		}
	}()


	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}