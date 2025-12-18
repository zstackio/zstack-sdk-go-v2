// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSAliyunSmsEndpointDetailParam CreateSNSAliyunSmsEndpoint detail param
type CreateSNSAliyunSmsEndpointDetailParam struct {
	AccessKeyUuid string `json:"accessKeyUuid" validate:"required"`
	Receivers []string `json:"receivers,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSAliyunSmsEndpointParam CreateSNSAliyunSmsEndpoint request param
type CreateSNSAliyunSmsEndpointParam struct {
	BaseParam
	Params CreateSNSAliyunSmsEndpointDetailParam `json:"params"`
}
