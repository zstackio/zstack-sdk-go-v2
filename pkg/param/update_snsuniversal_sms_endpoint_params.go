// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSUniversalSmsEndpointDetailParam UpdateSNSUniversalSmsEndpoint detail param
type UpdateSNSUniversalSmsEndpointDetailParam struct {
	SmsAccessKeyId string `json:"smsAccessKeyId" validate:"required"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSUniversalSmsEndpointParam UpdateSNSUniversalSmsEndpoint request param
type UpdateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params UpdateSNSUniversalSmsEndpointDetailParam `json:"params"`
}
