package exporter

import (
	"linkwan.cn/linkExporter/src/cache"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	expireTimeDuration = time.Duration(1) * time.Hour
	offlineDuration    = time.Duration(15) * time.Second
)

type channelExporter struct {
	sn           string
	tunnelMetric *prometheus.Desc
	cpeMetric    *prometheus.Desc

	reportTime time.Time
	reg        *prometheus.Registry
}

func (e *channelExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.tunnelMetric
	ch <- e.cpeMetric
}

func (e *channelExporter) Collect(ch chan<- prometheus.Metric) {

	channelStateCache := cache.GetChannelStateCache()
	if value, ok := channelStateCache.Load(e.sn); ok {
		channelMetricValue := 1
		cpeMetricValue := 1

		state := value.(*cache.ChannelState_t)

		if state.LastReportTime.Add(offlineDuration).Before(time.Now()) {
			cpeMetricValue = 0
			channelMetricValue = 0
		} else {
			if state.MasterTunnelState == cache.Down &&
				state.SlaveTunnelState == cache.Down {
				channelMetricValue = 0
			}
		}
		e.reportTime = state.LastReportTime

		ch <- prometheus.MustNewConstMetric(e.tunnelMetric, prometheus.CounterValue, float64(channelMetricValue))
		ch <- prometheus.MustNewConstMetric(e.cpeMetric, prometheus.CounterValue, float64(cpeMetricValue))

	} else {

	}

	// todo: unregister at expire time
	//if e.givenTime.Add(expireTimeDuration).Before(time.Now()) {
	//	e.reg.Unregister(e)
	//}
}

func NewChannelExporter(sn string, registry *prometheus.Registry) *channelExporter {
	return &channelExporter{
		sn:           sn,
		tunnelMetric: prometheus.NewDesc("tunnel_state", "", []string{}, prometheus.Labels{"SN": sn}),
		cpeMetric:    prometheus.NewDesc("cpe_state", "", []string{}, prometheus.Labels{"SN": sn}),
		reportTime:   time.Now(),
		//reg: registry,
	}
}
