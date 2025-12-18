// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsVSwitchDetailParam UpdateEcsVSwitch详细参数
type UpdateEcsVSwitchDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateEcsVSwitchParam UpdateEcsVSwitch请求参数
type UpdateEcsVSwitchParam struct {
	BaseParam
	Params UpdateEcsVSwitchDetailParam `json:"params"` // 详细参数
}

