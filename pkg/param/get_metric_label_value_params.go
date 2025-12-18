// Copyright (c) ZStack.io, Inc.

package param

// GetMetricLabelValueDetailParam GetMetricLabelValue detail param
type GetMetricLabelValueDetailParam struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	LabelNames []string `json:"labelNames" validate:"required"`
	FilterLabels []string `json:"filterLabels,omitempty"`
}

// GetMetricLabelValueParam GetMetricLabelValue request param
type GetMetricLabelValueParam struct {
	BaseParam
	Params GetMetricLabelValueDetailParam `json:"params"`
}
