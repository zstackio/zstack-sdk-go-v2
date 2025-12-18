// Copyright (c) ZStack.io, Inc.

package param

// DetachL3NetworksFromIPsecConnectionDetailParam DetachL3NetworksFromIPsecConnection detail param
type DetachL3NetworksFromIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
}

// DetachL3NetworksFromIPsecConnectionParam DetachL3NetworksFromIPsecConnection request param
type DetachL3NetworksFromIPsecConnectionParam struct {
	BaseParam
	Params DetachL3NetworksFromIPsecConnectionDetailParam `json:"params"`
}
