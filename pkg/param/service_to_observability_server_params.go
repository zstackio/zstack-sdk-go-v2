// Copyright (c) ZStack.io, Inc.

package param

// AttachServiceToObservabilityServerDetailParam AttachServiceToObservabilityServer详细参数
type AttachServiceToObservabilityServerDetailParam struct {
	rest string `json:"observabilityServerUuid" validate:"required"` // 必填
	rest string `json:"serviceType" validate:"required"` // 必填
	rest string `json:"serviceUuid" validate:"required"` // 必填
}

// AttachServiceToObservabilityServerParam AttachServiceToObservabilityServer请求参数
type AttachServiceToObservabilityServerParam struct {
	BaseParam
	Params AttachServiceToObservabilityServerDetailParam `json:"params"` // 详细参数
}

