package new_relic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"new-relic-exporter/config"
	"strconv"
	"strings"
	"time"
)

const (
	metricDataUrl = "https://api.newrelic.com/v2/applications/%d/metrics/data.json"
)

func RetrieveMetricData(metricNames, metricValueNames []string, from, to time.Time, appId int) ([]*MetricData, error) {
	var (
		resp     MetricDataResponse
		endpoint = fmt.Sprintf(metricDataUrl, appId)
		formData = url.Values{}
	)

	for _, metricName := range metricNames {
		formData.Add("names[]", metricName)
	}

	for _, metricValue := range metricValueNames {
		formData.Add("values[]", metricValue)
	}

	formData.Set("summary", "true")
	formData.Set("from", from.String())
	formData.Set("to", to.String())

	encodedFormData := formData.Encode()

	httpClient := &http.Client{}
	r, err := http.NewRequest("GET", endpoint, strings.NewReader(encodedFormData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(encodedFormData)))
	r.Header.Add("X-Api-Key", *config.ApiKey)

	res, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.MetricData.Metrics, nil
}
