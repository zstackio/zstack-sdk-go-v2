// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSFeiShuEndpointDetailParam UpdateSNSFeiShuEndpoint详细参数
type UpdateSNSFeiShuEndpointDetailParam struct {
	rest string `json:"url,omitempty"`
	rest bool `json:"atAll,omitempty"`
	rest string `json:"secret,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
}

// UpdateSNSFeiShuEndpointParam UpdateSNSFeiShuEndpoint请求参数
type UpdateSNSFeiShuEndpointParam struct {
	BaseParam
	Params UpdateSNSFeiShuEndpointDetailParam `json:"params"` // 详细参数
}

