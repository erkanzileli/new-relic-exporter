package metrics

import "github.com/prometheus/client_golang/prometheus"

// MetricVectors contains all metric vectors for an app
var MetricVectors map[string]*prometheus.GaugeVec

var (
	AverageResponseTimeGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "average_response_time",
	}, []string{"application_name", "metric_name"})

	RequestsPerMinuteGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "requests_per_minute",
	}, []string{"application_name", "metric_name"})

	ErrorsPerMinuteGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "errors_per_minute",
	}, []string{"application_name", "metric_name"})

	ApdexTargetGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apdex_target",
	}, []string{"application_name"})

	ApdexScoreGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apdex_score",
	}, []string{"application_name"})

	HostCountGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "host_count",
	}, []string{"application_name"})

	InstanceCountGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "instance_count",
	}, []string{"application_name"})

	ConcurrentInstanceCountGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "concurrent_instance_count",
	}, []string{"application_name"})
)

func EnsureMetricVectorsExists(metricValueNames []string) {
	for _, metricValue := range metricValueNames {

		if _, ok := MetricVectors[metricValue]; !ok {

			MetricVectors[metricValue] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: metricValue,
			}, []string{"application_name", "metric_name"})
			prometheus.MustRegister(MetricVectors[metricValue])
		}
	}
}
