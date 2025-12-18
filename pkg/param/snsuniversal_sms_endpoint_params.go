// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSUniversalSmsEndpointDetailParam UpdateSNSUniversalSmsEndpoint详细参数
type UpdateSNSUniversalSmsEndpointDetailParam struct {
	rest string `json:"smsAccessKeyId" validate:"required"` // 必填
	rest string `json:"smsAccessKeySecret" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
}

// UpdateSNSUniversalSmsEndpointParam UpdateSNSUniversalSmsEndpoint请求参数
type UpdateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params UpdateSNSUniversalSmsEndpointDetailParam `json:"params"` // 详细参数
}

