// Copyright (c) ZStack.io, Inc.

package param

// GetVSwitchTypesDetailParam GetVSwitchTypes详细参数
type GetVSwitchTypesDetailParam struct {
}

// GetVSwitchTypesParam GetVSwitchTypes请求参数
type GetVSwitchTypesParam struct {
	BaseParam
	Params GetVSwitchTypesDetailParam `json:"params"` // 详细参数
}

