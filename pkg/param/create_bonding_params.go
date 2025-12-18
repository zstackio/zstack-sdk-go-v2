// Copyright (c) ZStack.io, Inc.

package param

// CreateBondingDetailParam CreateBonding detail param
type CreateBondingDetailParam struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	BondingName string `json:"bondingName" validate:"required"`
	SlaveUuids []string `json:"slaveUuids,omitempty"`
	SlaveNames []string `json:"slaveNames,omitempty"`
	Type string `json:"type" validate:"required"`
	Mode string `json:"mode" validate:"required"`
	XmitHashPolicy string `json:"xmitHashPolicy,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBondingParam CreateBonding request param
type CreateBondingParam struct {
	BaseParam
	Params CreateBondingDetailParam `json:"params"`
}
