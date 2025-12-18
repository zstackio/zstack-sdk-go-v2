// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmConsolePasswordDetailParam DeleteVmConsolePassword detail param
type DeleteVmConsolePasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteVmConsolePasswordParam DeleteVmConsolePassword request param
type DeleteVmConsolePasswordParam struct {
	BaseParam
	Params DeleteVmConsolePasswordDetailParam `json:"params"`
}
