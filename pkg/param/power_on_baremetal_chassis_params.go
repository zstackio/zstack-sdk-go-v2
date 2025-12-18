// Copyright (c) ZStack.io, Inc.

package param

// PowerOnBaremetalChassisDetailParam PowerOnBaremetalChassis详细参数
type PowerOnBaremetalChassisDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
}

// PowerOnBaremetalChassisParam PowerOnBaremetalChassis请求参数
type PowerOnBaremetalChassisParam struct {
	BaseParam
	Params PowerOnBaremetalChassisDetailParam `json:"params"` // 详细参数
}

