// Copyright (c) ZStack.io, Inc.

package param

// GetPrometheusMetricLabelValueDetailParam GetPrometheusMetricLabelValue detail param
type GetPrometheusMetricLabelValueDetailParam struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	LabelNames []string `json:"labelNames,omitempty"`
	FilterLabels []string `json:"filterLabels,omitempty"`
}

// GetPrometheusMetricLabelValueParam GetPrometheusMetricLabelValue request param
type GetPrometheusMetricLabelValueParam struct {
	BaseParam
	Params GetPrometheusMetricLabelValueDetailParam `json:"params"`
}
