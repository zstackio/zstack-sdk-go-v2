// Copyright (c) ZStack.io, Inc.

package param

// DetachNetworkServiceFromL3NetworkDetailParam DetachNetworkServiceFromL3Network detail param
type DetachNetworkServiceFromL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	NetworkServices map[string]interface{} `json:"networkServices,omitempty"`
	Service string `json:"service,omitempty"`
}

// DetachNetworkServiceFromL3NetworkParam DetachNetworkServiceFromL3Network request param
type DetachNetworkServiceFromL3NetworkParam struct {
	BaseParam
	Params DetachNetworkServiceFromL3NetworkDetailParam `json:"params"`
}
