// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2IpmiChassisDetailParam UpdateBareMetal2IpmiChassis详细参数
type UpdateBareMetal2IpmiChassisDetailParam struct {
	rest string `json:"ipmiAddress,omitempty"`
	rest int `json:"ipmiPort,omitempty"`
	rest string `json:"ipmiUsername,omitempty"`
	rest string `json:"ipmiPassword,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateBareMetal2IpmiChassisParam UpdateBareMetal2IpmiChassis请求参数
type UpdateBareMetal2IpmiChassisParam struct {
	BaseParam
	Params UpdateBareMetal2IpmiChassisDetailParam `json:"params"` // 详细参数
}

