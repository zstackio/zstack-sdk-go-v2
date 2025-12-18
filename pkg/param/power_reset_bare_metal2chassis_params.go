// Copyright (c) ZStack.io, Inc.

package param

// PowerResetBareMetal2ChassisDetailParam PowerResetBareMetal2Chassis详细参数
type PowerResetBareMetal2ChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"bootDev,omitempty"`
}

// PowerResetBareMetal2ChassisParam PowerResetBareMetal2Chassis请求参数
type PowerResetBareMetal2ChassisParam struct {
	BaseParam
	Params PowerResetBareMetal2ChassisDetailParam `json:"params"` // 详细参数
}

