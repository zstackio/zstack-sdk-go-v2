// Copyright (c) ZStack.io, Inc.

package param

// SetServiceTypeOnHostNetworkBondingDetailParam SetServiceTypeOnHostNetworkBonding detail param
type SetServiceTypeOnHostNetworkBondingDetailParam struct {
	BondingUuids []string `json:"bondingUuids" validate:"required"`
	VlanIds []int `json:"vlanIds,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
}

// SetServiceTypeOnHostNetworkBondingParam SetServiceTypeOnHostNetworkBonding request param
type SetServiceTypeOnHostNetworkBondingParam struct {
	BaseParam
	Params SetServiceTypeOnHostNetworkBondingDetailParam `json:"params"`
}
