// Copyright (c) ZStack.io, Inc.

package param

// DeleteMetricDataDetailParam DeleteMetricData detail param
type DeleteMetricDataDetailParam struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	Labels []string `json:"labels,omitempty"`
}

// DeleteMetricDataParam DeleteMetricData request param
type DeleteMetricDataParam struct {
	BaseParam
	Params DeleteMetricDataDetailParam `json:"params"`
}
