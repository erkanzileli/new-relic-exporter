package new_relic

import (
	"encoding/json"
	"time"
)

type MetricDataResponse struct {
	MetricData struct {
		From            *time.Time    `json:"from"`
		To              *time.Time    `json:"to"`
		MetricsNotFound []string      `json:"metrics_not_found"`
		MetricsFound    []string      `json:"metrics_found"`
		Metrics         []*MetricData `json:"metrics"`
	} `json:"metric_data"`
}

type MetricData struct {
	Name       string            `json:"name,omitempty"`
	TimeSlices []MetricTimeSlice `json:"timeslices,omitempty"`
}

type MetricTimeSlice struct {
	From   *time.Time            `json:"from"`
	To     *time.Time            `json:"to"`
	Values MetricTimeSliceValues `json:"values"`
}

type MetricTimeSliceValues struct {
	AsPercentage           float64 `json:"as_percentage,omitempty"`
	AverageTime            float64 `json:"average_time,omitempty"`
	CallsPerMinute         float64 `json:"calls_per_minute,omitempty"`
	MaxValue               float64 `json:"max_value,omitempty"`
	TotalCallTimePerMinute float64 `json:"total_call_time_per_minute,omitempty"`
	Utilization            float64 `json:"utilization,omitempty"`

	Values map[string]float64 `json:"-"`
}

func (m *MetricTimeSliceValues) UnmarshalJSON(bytes []byte) error {
	type timeSliceValues MetricTimeSliceValues
	metricValues := timeSliceValues{
		Values: map[string]float64{},
	}

	if err := json.Unmarshal(bytes, &metricValues); err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &metricValues.Values); err != nil {
		return err
	}

	*m = MetricTimeSliceValues(metricValues)
	return nil
}
