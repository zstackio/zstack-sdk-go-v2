// Copyright (c) ZStack.io, Inc.

package param

// GetAllMetricMetadataDetailParam GetAllMetricMetadata详细参数
type GetAllMetricMetadataDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"namespace,omitempty"`
}

// GetAllMetricMetadataParam GetAllMetricMetadata请求参数
type GetAllMetricMetadataParam struct {
	BaseParam
	Params GetAllMetricMetadataDetailParam `json:"params"` // 详细参数
}

