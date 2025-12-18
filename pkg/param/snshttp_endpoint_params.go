// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSHttpEndpointDetailParam UpdateSNSHttpEndpoint详细参数
type UpdateSNSHttpEndpointDetailParam struct {
	rest string `json:"url,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
}

// UpdateSNSHttpEndpointParam UpdateSNSHttpEndpoint请求参数
type UpdateSNSHttpEndpointParam struct {
	BaseParam
	Params UpdateSNSHttpEndpointDetailParam `json:"params"` // 详细参数
}

