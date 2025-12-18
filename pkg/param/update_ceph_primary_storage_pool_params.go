// Copyright (c) ZStack.io, Inc.

package param

// UpdateCephPrimaryStoragePoolDetailParam UpdateCephPrimaryStoragePool detail param
type UpdateCephPrimaryStoragePoolDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AliasName string `json:"aliasName,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateCephPrimaryStoragePoolParam UpdateCephPrimaryStoragePool request param
type UpdateCephPrimaryStoragePoolParam struct {
	BaseParam
	Params UpdateCephPrimaryStoragePoolDetailParam `json:"params"`
}
