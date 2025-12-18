// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryPassThroughDetailParam PrometheusQueryPassThrough详细参数
type PrometheusQueryPassThroughDetailParam struct {
	rest bool `json:"instant,omitempty"`
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest string `json:"step,omitempty"`
	rest string `json:"expression" validate:"required"` // 必填
	rest string `json:"relativeTime,omitempty"`
}

// PrometheusQueryPassThroughParam PrometheusQueryPassThrough请求参数
type PrometheusQueryPassThroughParam struct {
	BaseParam
	Params PrometheusQueryPassThroughDetailParam `json:"params"` // 详细参数
}

