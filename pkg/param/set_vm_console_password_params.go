// Copyright (c) ZStack.io, Inc.

package param

// SetVmConsolePasswordDetailParam SetVmConsolePassword detail param
type SetVmConsolePasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ConsolePassword string `json:"consolePassword" validate:"required"`
}

// SetVmConsolePasswordParam SetVmConsolePassword request param
type SetVmConsolePasswordParam struct {
	BaseParam
	Params SetVmConsolePasswordDetailParam `json:"params"`
}
