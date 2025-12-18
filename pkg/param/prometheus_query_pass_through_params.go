// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryPassThroughDetailParam PrometheusQueryPassThrough detail param
type PrometheusQueryPassThroughDetailParam struct {
	Instant bool `json:"instant,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	Step string `json:"step,omitempty"`
	Expression string `json:"expression" validate:"required"`
	RelativeTime string `json:"relativeTime,omitempty"`
}

// PrometheusQueryPassThroughParam PrometheusQueryPassThrough request param
type PrometheusQueryPassThroughParam struct {
	BaseParam
	Params PrometheusQueryPassThroughDetailParam `json:"params"`
}
