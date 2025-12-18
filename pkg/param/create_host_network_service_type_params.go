// Copyright (c) ZStack.io, Inc.

package param

// CreateHostNetworkServiceTypeDetailParam CreateHostNetworkServiceType detail param
type CreateHostNetworkServiceTypeDetailParam struct {
	ServiceType string `json:"serviceType" validate:"required"`
	System bool `json:"system,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHostNetworkServiceTypeParam CreateHostNetworkServiceType request param
type CreateHostNetworkServiceTypeParam struct {
	BaseParam
	Params CreateHostNetworkServiceTypeDetailParam `json:"params"`
}
