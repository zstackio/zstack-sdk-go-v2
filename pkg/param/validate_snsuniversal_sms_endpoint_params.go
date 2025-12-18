// Copyright (c) ZStack.io, Inc.

package param

// ValidateSNSUniversalSmsEndpointDetailParam ValidateSNSUniversalSmsEndpoint详细参数
type ValidateSNSUniversalSmsEndpointDetailParam struct {
	rest string `json:"testMsg" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"phoneNumbers" validate:"required"` // 必填
}

// ValidateSNSUniversalSmsEndpointParam ValidateSNSUniversalSmsEndpoint请求参数
type ValidateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params ValidateSNSUniversalSmsEndpointDetailParam `json:"params"` // 详细参数
}

