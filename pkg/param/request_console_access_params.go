// Copyright (c) ZStack.io, Inc.

package param

// RequestConsoleAccessDetailParam RequestConsoleAccess detail param
type RequestConsoleAccessDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// RequestConsoleAccessParam RequestConsoleAccess request param
type RequestConsoleAccessParam struct {
	BaseParam
	Params RequestConsoleAccessDetailParam `json:"params"`
}
