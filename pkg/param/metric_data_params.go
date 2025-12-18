// Copyright (c) ZStack.io, Inc.

package param

// GetMetricDataDetailParam GetMetricData详细参数
type GetMetricDataDetailParam struct {
	rest string `json:"namespace" validate:"required"` // 必填
	rest string `json:"metricName" validate:"required"` // 必填
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest int64 `json:"offsetAheadOfCurrentTime,omitempty"`
	rest int `json:"period,omitempty"`
	rest []string `json:"labels,omitempty"`
	rest []string `json:"valueConditions,omitempty"`
	rest []string `json:"functions,omitempty"`
}

// GetMetricDataParam GetMetricData请求参数
type GetMetricDataParam struct {
	BaseParam
	Params GetMetricDataDetailParam `json:"params"` // 详细参数
}

