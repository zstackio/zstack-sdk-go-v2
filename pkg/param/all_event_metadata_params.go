// Copyright (c) ZStack.io, Inc.

package param

// GetAllEventMetadataDetailParam GetAllEventMetadata详细参数
type GetAllEventMetadataDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"namespace,omitempty"`
}

// GetAllEventMetadataParam GetAllEventMetadata请求参数
type GetAllEventMetadataParam struct {
	BaseParam
	Params GetAllEventMetadataDetailParam `json:"params"` // 详细参数
}

