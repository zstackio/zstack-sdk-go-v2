// Copyright (c) ZStack.io, Inc.

package param

// GetMetricDataDetailParam GetMetricData detail param
type GetMetricDataDetailParam struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	OffsetAheadOfCurrentTime int64 `json:"offsetAheadOfCurrentTime,omitempty"`
	Period int `json:"period,omitempty"`
	Labels []string `json:"labels,omitempty"`
	ValueConditions []string `json:"valueConditions,omitempty"`
	Functions []string `json:"functions,omitempty"`
}

// GetMetricDataParam GetMetricData request param
type GetMetricDataParam struct {
	BaseParam
	Params GetMetricDataDetailParam `json:"params"`
}
