// Copyright (c) ZStack.io, Inc.

package param

// AttachNetworkServiceToL3NetworkDetailParam AttachNetworkServiceToL3Network detail param
type AttachNetworkServiceToL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	NetworkServices map[string]interface{} `json:"networkServices" validate:"required"`
}

// AttachNetworkServiceToL3NetworkParam AttachNetworkServiceToL3Network request param
type AttachNetworkServiceToL3NetworkParam struct {
	BaseParam
	Params AttachNetworkServiceToL3NetworkDetailParam `json:"params"`
}
