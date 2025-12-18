// Copyright (c) ZStack.io, Inc.

package param

// GetVmConsoleAddressDetailParam GetVmConsoleAddress detail param
type GetVmConsoleAddressDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmConsoleAddressParam GetVmConsoleAddress request param
type GetVmConsoleAddressParam struct {
	BaseParam
	Params GetVmConsoleAddressDetailParam `json:"params"`
}
