// Copyright (c) ZStack.io, Inc.

package param

// AttachServiceToObservabilityServerDetailParam AttachServiceToObservabilityServer detail param
type AttachServiceToObservabilityServerDetailParam struct {
	ObservabilityServerUuid string `json:"observabilityServerUuid" validate:"required"`
	ServiceType string `json:"serviceType" validate:"required"`
	ServiceUuid string `json:"serviceUuid" validate:"required"`
}

// AttachServiceToObservabilityServerParam AttachServiceToObservabilityServer request param
type AttachServiceToObservabilityServerParam struct {
	BaseParam
	Params AttachServiceToObservabilityServerDetailParam `json:"params"`
}
