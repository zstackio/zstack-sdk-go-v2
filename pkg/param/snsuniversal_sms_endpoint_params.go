// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSNSUniversalSmsEndpointParamDetail UpdateSNSUniversalSmsEndpoint detail param
type UpdateSNSUniversalSmsEndpointParamDetail struct {
	SmsAccessKeyId string `json:"smsAccessKeyId" validate:"required"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
}

// UpdateSNSUniversalSmsEndpointParam UpdateSNSUniversalSmsEndpoint request param
type UpdateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params UpdateSNSUniversalSmsEndpointParamDetail `json:"updateSNSUniversalSmsEndpoint"`
}
// CreateSNSUniversalSmsEndpointParamDetail CreateSNSUniversalSmsEndpoint detail param
type CreateSNSUniversalSmsEndpointParamDetail struct {
	SmsAccessKeyId string `json:"smsAccessKeyId" validate:"required"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret" validate:"required"`
	Supplier string `json:"supplier" validate:"required"`
	AdditionParam map[string]string `json:"additionParam,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSUniversalSmsEndpointParam CreateSNSUniversalSmsEndpoint request param
type CreateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params CreateSNSUniversalSmsEndpointParamDetail `json:"params"`
}
// ValidateSNSUniversalSmsEndpointParamDetail ValidateSNSUniversalSmsEndpoint detail param
type ValidateSNSUniversalSmsEndpointParamDetail struct {
	TestMsg string `json:"testMsg" validate:"required"`
	PhoneNumbers []string `json:"phoneNumbers" validate:"required"`
}

// ValidateSNSUniversalSmsEndpointParam ValidateSNSUniversalSmsEndpoint request param
type ValidateSNSUniversalSmsEndpointParam struct {
	BaseParam
	Params ValidateSNSUniversalSmsEndpointParamDetail `json:"validateSNSUniversalSmsEndpoint"`
}
