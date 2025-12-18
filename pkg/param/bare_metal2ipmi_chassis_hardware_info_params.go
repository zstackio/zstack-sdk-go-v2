// Copyright (c) ZStack.io, Inc.

package param

// CreateBareMetal2IpmiChassisHardwareInfoDetailParam CreateBareMetal2IpmiChassisHardwareInfo详细参数
type CreateBareMetal2IpmiChassisHardwareInfoDetailParam struct {
	rest string `json:"ipmiAddress" validate:"required"` // 必填
	rest int `json:"ipmiPort" validate:"required"` // 必填
	rest string `json:"hardwareInfo" validate:"required"` // 必填
	rest string `json:"convertInfo,omitempty"`
}

// CreateBareMetal2IpmiChassisHardwareInfoParam CreateBareMetal2IpmiChassisHardwareInfo请求参数
type CreateBareMetal2IpmiChassisHardwareInfoParam struct {
	BaseParam
	Params CreateBareMetal2IpmiChassisHardwareInfoDetailParam `json:"params"` // 详细参数
}

