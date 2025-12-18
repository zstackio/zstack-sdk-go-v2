// Copyright (c) ZStack.io, Inc.

package param

// GetFactoryModeStateDetailParam GetFactoryModeState详细参数
type GetFactoryModeStateDetailParam struct {
}

// GetFactoryModeStateParam GetFactoryModeState请求参数
type GetFactoryModeStateParam struct {
	BaseParam
	Params GetFactoryModeStateDetailParam `json:"params"` // 详细参数
}

