// Copyright (c) ZStack.io, Inc.

package param

// ValidateSNSAliyunSmsEndpointDetailParam ValidateSNSAliyunSmsEndpoint详细参数
type ValidateSNSAliyunSmsEndpointDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"phoneNumbers" validate:"required"` // 必填
}

// ValidateSNSAliyunSmsEndpointParam ValidateSNSAliyunSmsEndpoint请求参数
type ValidateSNSAliyunSmsEndpointParam struct {
	BaseParam
	Params ValidateSNSAliyunSmsEndpointDetailParam `json:"params"` // 详细参数
}

