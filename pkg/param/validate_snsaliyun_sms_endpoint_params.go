// Copyright (c) ZStack.io, Inc.

package param

// ValidateSNSAliyunSmsEndpointDetailParam ValidateSNSAliyunSmsEndpoint detail param
type ValidateSNSAliyunSmsEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PhoneNumbers []string `json:"phoneNumbers" validate:"required"`
}

// ValidateSNSAliyunSmsEndpointParam ValidateSNSAliyunSmsEndpoint request param
type ValidateSNSAliyunSmsEndpointParam struct {
	BaseParam
	Params ValidateSNSAliyunSmsEndpointDetailParam `json:"params"`
}
