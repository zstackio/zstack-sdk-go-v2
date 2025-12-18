// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryMetadataDetailParam PrometheusQueryMetadata detail param
type PrometheusQueryMetadataDetailParam struct {
	Matches []string `json:"matches" validate:"required"`
}

// PrometheusQueryMetadataParam PrometheusQueryMetadata request param
type PrometheusQueryMetadataParam struct {
	BaseParam
	Params PrometheusQueryMetadataDetailParam `json:"params"`
}
