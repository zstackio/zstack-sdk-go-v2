// Copyright (c) ZStack.io, Inc.

package param

// UpdateBondingDetailParam UpdateBonding detail param
type UpdateBondingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SlaveUuids []string `json:"slaveUuids,omitempty"`
	SlaveNames []string `json:"slaveNames,omitempty"`
	Type string `json:"type,omitempty"`
	Mode string `json:"mode,omitempty"`
	XmitHashPolicy string `json:"xmitHashPolicy,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBondingParam UpdateBonding request param
type UpdateBondingParam struct {
	BaseParam
	Params UpdateBondingDetailParam `json:"params"`
}
