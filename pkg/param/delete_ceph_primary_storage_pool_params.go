// Copyright (c) ZStack.io, Inc.

package param

// DeleteCephPrimaryStoragePoolDetailParam DeleteCephPrimaryStoragePool detail param
type DeleteCephPrimaryStoragePoolDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteCephPrimaryStoragePoolParam DeleteCephPrimaryStoragePool request param
type DeleteCephPrimaryStoragePoolParam struct {
	BaseParam
	Params DeleteCephPrimaryStoragePoolDetailParam `json:"params"`
}
