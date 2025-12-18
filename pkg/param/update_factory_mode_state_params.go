// Copyright (c) ZStack.io, Inc.

package param

// UpdateFactoryModeStateDetailParam UpdateFactoryModeState detail param
type UpdateFactoryModeStateDetailParam struct {
	FactoryModeState bool `json:"factoryModeState" validate:"required"`
}

// UpdateFactoryModeStateParam UpdateFactoryModeState request param
type UpdateFactoryModeStateParam struct {
	BaseParam
	Params UpdateFactoryModeStateDetailParam `json:"params"`
}
