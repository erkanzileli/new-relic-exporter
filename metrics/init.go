package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	MetricVectors = map[string]*prometheus.GaugeVec{
		"average_response_time":     AverageResponseTimeGaugeVec,
		"requests_per_minute":       RequestsPerMinuteGaugeVec,
		"errors_per_minute":         ErrorsPerMinuteGaugeVec,
		"apdex_target":              ApdexTargetGaugeVec,
		"apdex_score":               ApdexScoreGaugeVec,
		"host_count":                HostCountGaugeVec,
		"instance_count":            InstanceCountGaugeVec,
		"concurrent_instance_count": ConcurrentInstanceCountGaugeVec,
	}

	prometheus.MustRegister(AverageResponseTimeGaugeVec)
	prometheus.MustRegister(RequestsPerMinuteGaugeVec)
	prometheus.MustRegister(ErrorsPerMinuteGaugeVec)
	prometheus.MustRegister(ApdexTargetGaugeVec)
	prometheus.MustRegister(ApdexScoreGaugeVec)
	prometheus.MustRegister(HostCountGaugeVec)
	prometheus.MustRegister(InstanceCountGaugeVec)
	prometheus.MustRegister(ConcurrentInstanceCountGaugeVec)
}
