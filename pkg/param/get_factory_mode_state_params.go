// Copyright (c) ZStack.io, Inc.

package param

// GetFactoryModeStateDetailParam GetFactoryModeState detail param
type GetFactoryModeStateDetailParam struct {
}

// GetFactoryModeStateParam GetFactoryModeState request param
type GetFactoryModeStateParam struct {
	BaseParam
	Params GetFactoryModeStateDetailParam `json:"params"`
}
