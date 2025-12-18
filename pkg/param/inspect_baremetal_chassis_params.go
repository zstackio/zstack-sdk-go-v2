// Copyright (c) ZStack.io, Inc.

package param

// InspectBaremetalChassisDetailParam InspectBaremetalChassis详细参数
type InspectBaremetalChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// InspectBaremetalChassisParam InspectBaremetalChassis请求参数
type InspectBaremetalChassisParam struct {
	BaseParam
	Params InspectBaremetalChassisDetailParam `json:"params"` // 详细参数
}

