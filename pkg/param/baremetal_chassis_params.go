// Copyright (c) ZStack.io, Inc.

package param

// UpdateBaremetalChassisDetailParam UpdateBaremetalChassis详细参数
type UpdateBaremetalChassisDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"ipmiAddress,omitempty"`
	rest int `json:"ipmiPort,omitempty"`
	rest string `json:"ipmiUsername,omitempty"`
	rest string `json:"ipmiPassword,omitempty"`
}

// UpdateBaremetalChassisParam UpdateBaremetalChassis请求参数
type UpdateBaremetalChassisParam struct {
	BaseParam
	Params UpdateBaremetalChassisDetailParam `json:"params"` // 详细参数
}

