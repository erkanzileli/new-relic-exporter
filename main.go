package main

import (
	"github.com/newrelic/newrelic-client-go/newrelic"
	"github.com/newrelic/newrelic-client-go/pkg/apm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"new-relic-exporter/config"
	"new-relic-exporter/metrics"
	"new-relic-exporter/new-relic"
	"new-relic-exporter/util"
	"regexp"
	"time"
)

func main() {
	var (
		client *newrelic.NewRelic
		err    error
	)

	if *config.PersonalApiKey {
		client, err = newrelic.New(newrelic.ConfigPersonalAPIKey(*config.ApiKey))
	} else {
		client, err = newrelic.New(newrelic.ConfigAdminAPIKey(*config.ApiKey))
	}

	if err != nil {
		log.Fatal("error initializing client:", err)
	}

	c := cron.New()
	c.AddFunc(*config.ScrapeInterval, func() {

		log.Println("Started for", *config.AppId, "as goroutineId:", util.GetCallerName())

		application, err := client.APM.GetApplication(*config.AppId)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Application Name:", application.Name)

		fillApplicationSummary(application)

		to := time.Now()
		from := to.Add(-1 * time.Minute)
		metricNames, metricValueNames := retrieveMetricNamesAndValues(client, *config.AppId)

		metrics.EnsureMetricVectorsExists(metricValueNames)

		metricData, err := new_relic.RetrieveMetricData(metricNames, metricValueNames, from, to, *config.AppId)
		if err != nil {
			log.Fatal(err)
		}

		fillMetricData(metricData, application, metricValueNames)

		log.Println("Done for", application.Name)

	})

	c.Start()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*config.Addr, nil))
}

func retrieveMetricNamesAndValues(client *newrelic.NewRelic, appId int) ([]string, []string) {
	metrics, err := client.APM.GetMetricNames(appId, apm.MetricNamesParams{})
	if err != nil {
		log.Fatal(err)
	}

	filterRegex := regexp.MustCompile("^(WebTransaction|Errors|Datastore|Memory)/.*$")

	metricNames := make([]string, 0)
	metricValues := make([]string, 0)
	metricNameSet := make(map[string]bool)
	metricValueSet := make(map[string]bool)

	for _, metric := range metrics {
		if ok := metricNameSet[metric.Name]; ok {
			continue
		}

		if !filterRegex.MatchString(metric.Name) {
			continue
		}

		metricNames = append(metricNames, metric.Name)
		metricNameSet[metric.Name] = true

		for _, metricValue := range metric.Values {
			if len(metricValue) == 0 {
				continue
			}

			if ok := metricValueSet[metricValue]; ok {
				continue
			}

			metricValues = append(metricValues, metricValue)
			metricValueSet[metricValue] = true
		}
	}

	return metricNames, metricValues
}

func fillMetricData(metricData []*new_relic.MetricData, application *apm.Application, metricValueNames []string) {
	for _, datum := range metricData {
		for _, metricValueName := range metricValueNames {
			if value, ok := datum.TimeSlices[0].Values.Values[metricValueName]; ok {
				if vector, ok := metrics.MetricVectors[metricValueName]; ok {
					vector.
						With(prometheus.Labels{"application_name": application.Name, "metric_name": datum.Name}).
						Set(value)
				}
			}
		}
	}
}

func fillApplicationSummary(application *apm.Application) {
	metrics.AverageResponseTimeGaugeVec.
		With(prometheus.Labels{"application_name": application.Name, "metric_name": "application_summary"}).
		Set(application.Summary.ResponseTime)

	metrics.RequestsPerMinuteGaugeVec.
		With(prometheus.Labels{"application_name": application.Name, "metric_name": "application_summary"}).
		Set(application.Summary.Throughput)

	metrics.ErrorsPerMinuteGaugeVec.
		With(prometheus.Labels{"application_name": application.Name, "metric_name": "application_summary"}).
		Set(application.Summary.ErrorRate)

	metrics.ApdexTargetGaugeVec.
		With(prometheus.Labels{"application_name": application.Name}).
		Set(application.Summary.ApdexTarget)

	metrics.ApdexScoreGaugeVec.
		With(prometheus.Labels{"application_name": application.Name}).
		Set(application.Summary.ApdexScore)

	metrics.HostCountGaugeVec.
		With(prometheus.Labels{"application_name": application.Name}).
		Set(float64(application.Summary.HostCount))

	metrics.InstanceCountGaugeVec.
		With(prometheus.Labels{"application_name": application.Name}).
		Set(float64(application.Summary.InstanceCount))

	metrics.ConcurrentInstanceCountGaugeVec.
		With(prometheus.Labels{"application_name": application.Name}).
		Set(float64(application.Summary.ConcurrentInstanceCount))
}
