// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmCdRomDetailParam UpdateVmCdRom detail param
type UpdateVmCdRomDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateVmCdRomParam UpdateVmCdRom request param
type UpdateVmCdRomParam struct {
	BaseParam
	Params UpdateVmCdRomDetailParam `json:"params"`
}
