// Copyright (c) ZStack.io, Inc.

package param

// DebugSignalDetailParam DebugSignal detail param
type DebugSignalDetailParam struct {
	Signals []string `json:"signals" validate:"required"`
}

// DebugSignalParam DebugSignal request param
type DebugSignalParam struct {
	BaseParam
	Params DebugSignalDetailParam `json:"params"`
}
