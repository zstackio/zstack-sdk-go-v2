// Copyright (c) ZStack.io, Inc.

package param

// ChangeL2NetworkVlanIdDetailParam ChangeL2NetworkVlanId detail param
type ChangeL2NetworkVlanIdDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Vlan int `json:"vlan,omitempty"`
	Type string `json:"type,omitempty"`
}

// ChangeL2NetworkVlanIdParam ChangeL2NetworkVlanId request param
type ChangeL2NetworkVlanIdParam struct {
	BaseParam
	Params ChangeL2NetworkVlanIdDetailParam `json:"params"`
}
