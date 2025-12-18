// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSEmailEndpointDetailParam CreateSNSEmailEndpoint详细参数
type CreateSNSEmailEndpointDetailParam struct {
	rest string `json:"email,omitempty"`
	rest []string `json:"emails,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSNSEmailEndpointParam CreateSNSEmailEndpoint请求参数
type CreateSNSEmailEndpointParam struct {
	BaseParam
	Params CreateSNSEmailEndpointDetailParam `json:"params"` // 详细参数
}

