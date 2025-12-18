// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmCdRomDetailParam DeleteVmCdRom detail param
type DeleteVmCdRomDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmCdRomParam DeleteVmCdRom request param
type DeleteVmCdRomParam struct {
	BaseParam
	Params DeleteVmCdRomDetailParam `json:"params"`
}
