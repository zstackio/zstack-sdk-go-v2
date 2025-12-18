// Copyright (c) ZStack.io, Inc.

package param

// PowerOnBareMetal2ChassisDetailParam PowerOnBareMetal2Chassis详细参数
type PowerOnBareMetal2ChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"bootDev,omitempty"`
}

// PowerOnBareMetal2ChassisParam PowerOnBareMetal2Chassis请求参数
type PowerOnBareMetal2ChassisParam struct {
	BaseParam
	Params PowerOnBareMetal2ChassisDetailParam `json:"params"` // 详细参数
}

