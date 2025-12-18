// Copyright (c) ZStack.io, Inc.

package param

// DetachServiceFromObservabilityServerDetailParam DetachServiceFromObservabilityServer详细参数
type DetachServiceFromObservabilityServerDetailParam struct {
	rest string `json:"observabilityServerUuid" validate:"required"` // 必填
	rest string `json:"serviceType" validate:"required"` // 必填
	rest string `json:"serviceUuid" validate:"required"` // 必填
}

// DetachServiceFromObservabilityServerParam DetachServiceFromObservabilityServer请求参数
type DetachServiceFromObservabilityServerParam struct {
	BaseParam
	Params DetachServiceFromObservabilityServerDetailParam `json:"params"` // 详细参数
}

