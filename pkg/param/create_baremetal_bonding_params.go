// Copyright (c) ZStack.io, Inc.

package param

// CreateBaremetalBondingDetailParam CreateBaremetalBonding detail param
type CreateBaremetalBondingDetailParam struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Mode int `json:"mode" validate:"required"`
	Slaves string `json:"slaves" validate:"required"`
	Opts string `json:"opts,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalBondingParam CreateBaremetalBonding request param
type CreateBaremetalBondingParam struct {
	BaseParam
	Params CreateBaremetalBondingDetailParam `json:"params"`
}
