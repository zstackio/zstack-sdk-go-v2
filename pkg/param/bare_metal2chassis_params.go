// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2ChassisDetailParam UpdateBareMetal2Chassis详细参数
type UpdateBareMetal2ChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisParam UpdateBareMetal2Chassis请求参数
type UpdateBareMetal2ChassisParam struct {
	BaseParam
	Params UpdateBareMetal2ChassisDetailParam `json:"params"` // 详细参数
}

