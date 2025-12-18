// Copyright (c) ZStack.io, Inc.

package param

// DeleteIpAddressDetailParam DeleteIpAddress detail param
type DeleteIpAddressDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	UsedIpUuids []string `json:"usedIpUuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIpAddressParam DeleteIpAddress request param
type DeleteIpAddressParam struct {
	BaseParam
	Params DeleteIpAddressDetailParam `json:"params"`
}
