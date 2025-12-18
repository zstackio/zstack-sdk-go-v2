// Copyright (c) ZStack.io, Inc.

package param

// SetVmInstanceDefaultCdRomDetailParam SetVmInstanceDefaultCdRom detail param
type SetVmInstanceDefaultCdRomDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// SetVmInstanceDefaultCdRomParam SetVmInstanceDefaultCdRom request param
type SetVmInstanceDefaultCdRomParam struct {
	BaseParam
	Params SetVmInstanceDefaultCdRomDetailParam `json:"params"`
}
