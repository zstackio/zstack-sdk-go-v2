// Copyright (c) ZStack.io, Inc.

package param

// GetAvailableTriggersDetailParam GetAvailableTriggers detail param
type GetAvailableTriggersDetailParam struct {
}

// GetAvailableTriggersParam GetAvailableTriggers request param
type GetAvailableTriggersParam struct {
	BaseParam
	Params GetAvailableTriggersDetailParam `json:"params"`
}
