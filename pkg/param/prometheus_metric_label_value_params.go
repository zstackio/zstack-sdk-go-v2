// Copyright (c) ZStack.io, Inc.

package param

// GetPrometheusMetricLabelValueDetailParam GetPrometheusMetricLabelValue详细参数
type GetPrometheusMetricLabelValueDetailParam struct {
	rest string `json:"namespace" validate:"required"` // 必填
	rest string `json:"metricName" validate:"required"` // 必填
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest []string `json:"labelNames,omitempty"`
	rest []string `json:"filterLabels,omitempty"`
}

// GetPrometheusMetricLabelValueParam GetPrometheusMetricLabelValue请求参数
type GetPrometheusMetricLabelValueParam struct {
	BaseParam
	Params GetPrometheusMetricLabelValueDetailParam `json:"params"` // 详细参数
}

