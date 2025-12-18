// Copyright (c) ZStack.io, Inc.

package param

// GetVmConsolePasswordDetailParam GetVmConsolePassword detail param
type GetVmConsolePasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmConsolePasswordParam GetVmConsolePassword request param
type GetVmConsolePasswordParam struct {
	BaseParam
	Params GetVmConsolePasswordDetailParam `json:"params"`
}
