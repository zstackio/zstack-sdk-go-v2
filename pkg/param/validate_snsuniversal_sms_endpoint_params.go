// Copyright (c) ZStack.io, Inc.

package param

// ValidateSNSUniversalSmsEndpointDetailParam ValidateSNSUniversalSmsEndpoint detail param
type ValidateSNSUniversalSmsEndpointDetailParam struct {
	TestMsg string `json:"testMsg" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	PhoneNumbers []string `json:"phoneNumbers" validate:"required"`
}

// ValidateSNSUniversalSmsEndpointParam ValidateSNSUniversalSmsEndpoint request param
type ValidateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params ValidateSNSUniversalSmsEndpointDetailParam `json:"params"`
}
