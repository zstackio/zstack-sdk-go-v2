// Copyright (c) ZStack.io, Inc.

package param

// GetVmQgaDetailParam GetVmQga detail param
type GetVmQgaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmQgaParam GetVmQga request param
type GetVmQgaParam struct {
	BaseParam
	Params GetVmQgaDetailParam `json:"params"`
}
