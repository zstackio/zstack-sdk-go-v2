// Copyright (c) ZStack.io, Inc.

package param

// GetVmConsoleAddressDetailParam GetVmConsoleAddress详细参数
type GetVmConsoleAddressDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmConsoleAddressParam GetVmConsoleAddress请求参数
type GetVmConsoleAddressParam struct {
	BaseParam
	Params GetVmConsoleAddressDetailParam `json:"params"` // 详细参数
}

