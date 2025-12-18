// Copyright (c) ZStack.io, Inc.

package param

// UpdateConsolePasswordDetailParam UpdateConsolePassword detail param
type UpdateConsolePasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UpdateConsolePasswordParam UpdateConsolePassword request param
type UpdateConsolePasswordParam struct {
	BaseParam
	Params UpdateConsolePasswordDetailParam `json:"params"`
}
