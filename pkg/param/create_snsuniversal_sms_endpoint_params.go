// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSUniversalSmsEndpointDetailParam CreateSNSUniversalSmsEndpoint detail param
type CreateSNSUniversalSmsEndpointDetailParam struct {
	SmsAccessKeyId string `json:"smsAccessKeyId" validate:"required"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret" validate:"required"`
	Supplier string `json:"supplier" validate:"required"`
	AdditionParam map[string]string `json:"additionParam,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSUniversalSmsEndpointParam CreateSNSUniversalSmsEndpoint request param
type CreateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params CreateSNSUniversalSmsEndpointDetailParam `json:"params"`
}
