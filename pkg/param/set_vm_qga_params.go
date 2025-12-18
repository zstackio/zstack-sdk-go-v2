// Copyright (c) ZStack.io, Inc.

package param

// SetVmQgaDetailParam SetVmQga detail param
type SetVmQgaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Enable bool `json:"enable" validate:"required"`
}

// SetVmQgaParam SetVmQga request param
type SetVmQgaParam struct {
	BaseParam
	Params SetVmQgaDetailParam `json:"params"`
}
