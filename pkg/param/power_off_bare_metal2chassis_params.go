// Copyright (c) ZStack.io, Inc.

package param

// PowerOffBareMetal2ChassisDetailParam PowerOffBareMetal2Chassis详细参数
type PowerOffBareMetal2ChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// PowerOffBareMetal2ChassisParam PowerOffBareMetal2Chassis请求参数
type PowerOffBareMetal2ChassisParam struct {
	BaseParam
	Params PowerOffBareMetal2ChassisDetailParam `json:"params"` // 详细参数
}

