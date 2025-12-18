// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryLabelValuesDetailParam PrometheusQueryLabelValues detail param
type PrometheusQueryLabelValuesDetailParam struct {
	Labels []string `json:"labels" validate:"required"`
}

// PrometheusQueryLabelValuesParam PrometheusQueryLabelValues request param
type PrometheusQueryLabelValuesParam struct {
	BaseParam
	Params PrometheusQueryLabelValuesDetailParam `json:"params"`
}
