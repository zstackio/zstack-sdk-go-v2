// Copyright (c) ZStack.io, Inc.

package param

// PowerResetBaremetalChassisDetailParam PowerResetBaremetalChassis详细参数
type PowerResetBaremetalChassisDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
}

// PowerResetBaremetalChassisParam PowerResetBaremetalChassis请求参数
type PowerResetBaremetalChassisParam struct {
	BaseParam
	Params PowerResetBaremetalChassisDetailParam `json:"params"` // 详细参数
}

