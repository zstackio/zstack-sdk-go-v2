// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateSNSAliyunSmsEndpointParamDetail CreateSNSAliyunSmsEndpoint detail param
type CreateSNSAliyunSmsEndpointParamDetail struct {
	AccessKeyUuid string `json:"accessKeyUuid" validate:"required"`
	Receivers []string `json:"receivers,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSAliyunSmsEndpointParam CreateSNSAliyunSmsEndpoint request param
type CreateSNSAliyunSmsEndpointParam struct {
	BaseParam
	Params CreateSNSAliyunSmsEndpointParamDetail `json:"params"`
}
// ValidateSNSAliyunSmsEndpointParamDetail ValidateSNSAliyunSmsEndpoint detail param
type ValidateSNSAliyunSmsEndpointParamDetail struct {
	PhoneNumbers []string `json:"phoneNumbers" validate:"required"`
}

// ValidateSNSAliyunSmsEndpointParam ValidateSNSAliyunSmsEndpoint request param
type ValidateSNSAliyunSmsEndpointParam struct {
	BaseParam
	Params ValidateSNSAliyunSmsEndpointParamDetail `json:"validateSNSAliyunSmsEndpoint"`
}
