// Copyright (c) ZStack.io, Inc.

package param

// DetachServiceFromObservabilityServerDetailParam DetachServiceFromObservabilityServer detail param
type DetachServiceFromObservabilityServerDetailParam struct {
	ObservabilityServerUuid string `json:"observabilityServerUuid" validate:"required"`
	ServiceType string `json:"serviceType" validate:"required"`
	ServiceUuid string `json:"serviceUuid" validate:"required"`
}

// DetachServiceFromObservabilityServerParam DetachServiceFromObservabilityServer request param
type DetachServiceFromObservabilityServerParam struct {
	BaseParam
	Params DetachServiceFromObservabilityServerDetailParam `json:"params"`
}
