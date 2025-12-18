// Copyright (c) ZStack.io, Inc.

package param

// GetVmNumaDetailParam GetVmNuma detail param
type GetVmNumaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmNumaParam GetVmNuma request param
type GetVmNumaParam struct {
	BaseParam
	Params GetVmNumaDetailParam `json:"params"`
}
