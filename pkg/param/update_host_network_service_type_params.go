// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostNetworkServiceTypeDetailParam UpdateHostNetworkServiceType detail param
type UpdateHostNetworkServiceTypeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ServiceType string `json:"serviceType" validate:"required"`
	System bool `json:"system,omitempty"`
}

// UpdateHostNetworkServiceTypeParam UpdateHostNetworkServiceType request param
type UpdateHostNetworkServiceTypeParam struct {
	BaseParam
	Params UpdateHostNetworkServiceTypeDetailParam `json:"params"`
}
