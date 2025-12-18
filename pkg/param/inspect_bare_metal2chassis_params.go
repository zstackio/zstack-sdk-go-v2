// Copyright (c) ZStack.io, Inc.

package param

// InspectBareMetal2ChassisDetailParam InspectBareMetal2Chassis详细参数
type InspectBareMetal2ChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// InspectBareMetal2ChassisParam InspectBareMetal2Chassis请求参数
type InspectBareMetal2ChassisParam struct {
	BaseParam
	Params InspectBareMetal2ChassisDetailParam `json:"params"` // 详细参数
}

