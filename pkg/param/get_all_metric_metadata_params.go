// Copyright (c) ZStack.io, Inc.

package param

// GetAllMetricMetadataDetailParam GetAllMetricMetadata detail param
type GetAllMetricMetadataDetailParam struct {
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// GetAllMetricMetadataParam GetAllMetricMetadata request param
type GetAllMetricMetadataParam struct {
	BaseParam
	Params GetAllMetricMetadataDetailParam `json:"params"`
}
