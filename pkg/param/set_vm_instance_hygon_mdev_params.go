// Copyright (c) ZStack.io, Inc.

package param

// SetVmInstanceHygonMdevDetailParam SetVmInstanceHygonMdev detail param
type SetVmInstanceHygonMdevDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	HygonSecurityElementEnable string `json:"hygonSecurityElementEnable" validate:"required"`
}

// SetVmInstanceHygonMdevParam SetVmInstanceHygonMdev request param
type SetVmInstanceHygonMdevParam struct {
	BaseParam
	Params SetVmInstanceHygonMdevDetailParam `json:"params"`
}
