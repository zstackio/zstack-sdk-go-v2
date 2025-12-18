// Copyright (c) ZStack.io, Inc.

package param

// SetVmConsoleModeDetailParam SetVmConsoleMode detail param
type SetVmConsoleModeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Mode string `json:"mode" validate:"required"`
}

// SetVmConsoleModeParam SetVmConsoleMode request param
type SetVmConsoleModeParam struct {
	BaseParam
	Params SetVmConsoleModeDetailParam `json:"params"`
}
