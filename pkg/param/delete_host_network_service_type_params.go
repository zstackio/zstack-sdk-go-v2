// Copyright (c) ZStack.io, Inc.

package param

// DeleteHostNetworkServiceTypeDetailParam DeleteHostNetworkServiceType detail param
type DeleteHostNetworkServiceTypeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHostNetworkServiceTypeParam DeleteHostNetworkServiceType request param
type DeleteHostNetworkServiceTypeParam struct {
	BaseParam
	Params DeleteHostNetworkServiceTypeDetailParam `json:"params"`
}
