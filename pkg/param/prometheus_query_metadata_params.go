// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryMetadataDetailParam PrometheusQueryMetadata详细参数
type PrometheusQueryMetadataDetailParam struct {
	rest []string `json:"matches" validate:"required"` // 必填
}

// PrometheusQueryMetadataParam PrometheusQueryMetadata请求参数
type PrometheusQueryMetadataParam struct {
	BaseParam
	Params PrometheusQueryMetadataDetailParam `json:"params"` // 详细参数
}

