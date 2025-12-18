// Copyright (c) ZStack.io, Inc.

package param

// GetDebugSignalDetailParam GetDebugSignal detail param
type GetDebugSignalDetailParam struct {
}

// GetDebugSignalParam GetDebugSignal request param
type GetDebugSignalParam struct {
	BaseParam
	Params GetDebugSignalDetailParam `json:"params"`
}
