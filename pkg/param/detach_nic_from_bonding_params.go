// Copyright (c) ZStack.io, Inc.

package param

// DetachNicFromBondingDetailParam DetachNicFromBonding detail param
type DetachNicFromBondingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SlaveUuids []string `json:"slaveUuids" validate:"required"`
	Type string `json:"type,omitempty"`
}

// DetachNicFromBondingParam DetachNicFromBonding request param
type DetachNicFromBondingParam struct {
	BaseParam
	Params DetachNicFromBondingDetailParam `json:"params"`
}
