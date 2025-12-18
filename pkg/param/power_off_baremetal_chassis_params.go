// Copyright (c) ZStack.io, Inc.

package param

// PowerOffBaremetalChassisDetailParam PowerOffBaremetalChassis详细参数
type PowerOffBaremetalChassisDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
}

// PowerOffBaremetalChassisParam PowerOffBaremetalChassis请求参数
type PowerOffBaremetalChassisParam struct {
	BaseParam
	Params PowerOffBaremetalChassisDetailParam `json:"params"` // 详细参数
}

