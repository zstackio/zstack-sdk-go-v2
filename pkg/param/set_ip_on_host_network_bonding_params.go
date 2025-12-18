// Copyright (c) ZStack.io, Inc.

package param

// SetIpOnHostNetworkBondingDetailParam SetIpOnHostNetworkBonding detail param
type SetIpOnHostNetworkBondingDetailParam struct {
	BondingUuid string `json:"bondingUuid" validate:"required"`
	IpAddress string `json:"ipAddress,omitempty"`
	Netmask string `json:"netmask,omitempty"`
}

// SetIpOnHostNetworkBondingParam SetIpOnHostNetworkBonding request param
type SetIpOnHostNetworkBondingParam struct {
	BaseParam
	Params SetIpOnHostNetworkBondingDetailParam `json:"params"`
}
