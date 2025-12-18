// Copyright (c) ZStack.io, Inc.

package param

// AttachL3NetworksToIPsecConnectionDetailParam AttachL3NetworksToIPsecConnection detail param
type AttachL3NetworksToIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AttachL3NetworksToIPsecConnectionParam AttachL3NetworksToIPsecConnection request param
type AttachL3NetworksToIPsecConnectionParam struct {
	BaseParam
	Params AttachL3NetworksToIPsecConnectionDetailParam `json:"params"`
}
