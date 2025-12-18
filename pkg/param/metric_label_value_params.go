// Copyright (c) ZStack.io, Inc.

package param

// GetMetricLabelValueDetailParam GetMetricLabelValue详细参数
type GetMetricLabelValueDetailParam struct {
	rest string `json:"namespace" validate:"required"` // 必填
	rest string `json:"metricName" validate:"required"` // 必填
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest []string `json:"labelNames" validate:"required"` // 必填
	rest []string `json:"filterLabels,omitempty"`
}

// GetMetricLabelValueParam GetMetricLabelValue请求参数
type GetMetricLabelValueParam struct {
	BaseParam
	Params GetMetricLabelValueDetailParam `json:"params"` // 详细参数
}

