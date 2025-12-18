// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryLabelValuesDetailParam PrometheusQueryLabelValues详细参数
type PrometheusQueryLabelValuesDetailParam struct {
	rest []string `json:"labels" validate:"required"` // 必填
}

// PrometheusQueryLabelValuesParam PrometheusQueryLabelValues请求参数
type PrometheusQueryLabelValuesParam struct {
	BaseParam
	Params PrometheusQueryLabelValuesDetailParam `json:"params"` // 详细参数
}

