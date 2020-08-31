package exporter

import (
	"errors"
	"fmt"
	"linkwan.cn/linkExporter/src/config"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

var (
	registry *prometheus.Registry
)

func init() {
	registry = prometheus.NewPedanticRegistry()

	var parser expfmt.TextParser
	var parserText = func() ([]*dto.MetricFamily, error) {
		parsed, err := parser.TextToMetricFamilies(strings.NewReader(""))
		if err != nil {
			return nil, err
		}
		var result []*dto.MetricFamily
		for _, mf := range parsed {
			result = append(result, mf)
		}
		return result, nil
	}

	newGatherers := prometheus.Gatherers{
		//prometheus.DefaultGatherer,
		prometheus.GathererFunc(parserText),
		registry,
	}
	h := promhttp.HandlerFor(newGatherers, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { h.ServeHTTP(w, r) })
	fmt.Printf("Start Server at :%v\n", config.GetConfig().Port())
	p := fmt.Sprintf(":%s", config.GetConfig().Port())
	//if err := http.ListenAndServe(p, nil); err != nil {
	//	fmt.Printf("Error occur when start server %v", err)
	//	os.Exit(1)
	//}
	go http.ListenAndServe(p, nil)
}

func Register(sn string) error {
	if registry != nil {
		collector := NewChannelExporter(sn, registry)
		return registry.Register(collector)
	}
	return errors.New("no registry")
}
