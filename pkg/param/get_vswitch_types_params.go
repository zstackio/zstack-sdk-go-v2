// Copyright (c) ZStack.io, Inc.

package param

// GetVSwitchTypesDetailParam GetVSwitchTypes detail param
type GetVSwitchTypesDetailParam struct {
}

// GetVSwitchTypesParam GetVSwitchTypes request param
type GetVSwitchTypesParam struct {
	BaseParam
	Params GetVSwitchTypesDetailParam `json:"params"`
}
