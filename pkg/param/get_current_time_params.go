// Copyright (c) ZStack.io, Inc.

package param

// GetCurrentTimeDetailParam GetCurrentTime detail param
type GetCurrentTimeDetailParam struct {
}

// GetCurrentTimeParam GetCurrentTime request param
type GetCurrentTimeParam struct {
	BaseParam
	Params GetCurrentTimeDetailParam `json:"params"`
}
